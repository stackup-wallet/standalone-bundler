import { Client } from "userop";
import { TestAccount } from "../src/testAccount";
import config from "../config";
import { ethers } from "ethers";

describe("During the verification phase", () => {
  let client: Client;
  let acc: TestAccount;
  beforeAll(async () => {
    client = await Client.init(config.nodeUrl, {
      overrideBundlerRpc: config.bundlerUrl,
    });
    client.waitTimeoutMs = 2000;
    client.waitIntervalMs = 100;
    acc = await TestAccount.init(config.testAccount, config.nodeUrl, {
      overrideBundlerRpc: config.bundlerUrl,
    });
  });

  describe("With increasing call stack size", () => {
    describe("With zero sibling stacks", () => {
      [0, 2, 4, 6, 8, 10].forEach((depth) => {
        test(`Sender can run verification with ${depth} recursive calls`, async () => {
          const response = await client.sendUserOperation(
            acc.recursiveCall(depth, 0, 0)
          );
          const event = await response.wait();

          expect(event?.args.success).toBe(true);
        });
      });
    });
  });

  describe("With dependency on callGasLimit", () => {
    [0, 1, 2, 3, 4, 5].forEach((times) => {
      test(`Sender can run validation with non-simulated code that uses ${times} storage writes`, async () => {
        const response = await client.sendUserOperation(
          acc.forceValidationOOG(times)
        );
        const event = await response.wait();

        expect(event?.args.success).toBe(true);
      });
    });
  });

  describe("With increasing paymaster postOp gas used", () => {
    [0, 1, 2, 3, 4, 5].forEach((times) => {
      test(`Sender can run validation with non-simulated code that uses ${times} storage writes`, async () => {
        const response = await client.sendUserOperation(
          acc.forcePostOpValidationOOG(times),
          {
            onBuild(op) {
              console.log(
                ethers.BigNumber.from(op.verificationGasLimit).toNumber()
              );
              console.log(op.paymasterAndData);
            },
          }
        );
        const event = await response.wait();

        expect(event?.args.success).toBe(true);
      });
    });
  });
});
