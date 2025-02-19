## `generate_withdrawal_requests` Task

### Description
The `generate_withdrawal_requests` task facilitates the generation and submission of withdrawal requests to the Ethereum network. This task is crucial for simulating the process of withdrawing funds or exiting validators as part of execution layer triggered staking operations.

The source validators can be specified in three ways:
- By providing `sourcePubkey`, which is used as static pubkey for all requests
- By providing `sourceMnemonic`, `sourceStartIndex` & `sourceIndexCount` to select the source validators by the pubkeys derived from the mnemonic & key range
- By providing `sourceStartValidatorIndex` & `sourceIndexCount` to select the source validators by their validator index

### Configuration Parameters

- **`limitPerSlot`**:
  Specifies the maximum number of withdrawal requests allowed per slot, managing network load and ensuring efficient processing.

- **`limitTotal`**:
  Sets an upper limit on the total number of withdrawal requests that can be generated by this task.

- **`limitPending`**:
  Defines the maximum number of pending withdrawal requests allowed at any given time.

- **`sourcePubkey`**:
  The static pubkey to include in the withdrawal requests.

- **`sourceMnemonic`**:
  The mnemonic used to derive source validator keys from which withdrawals will be requested.

- **`sourceStartIndex`**:
  The starting index for key derivation from the source mnemonic.

- **`sourceStartValidatorIndex`**:
  An alternative to using `sourceMnemonic` and `sourceStartIndex`, this directly specifies the starting validator index for withdrawal requests.

- **`sourceIndexCount`**:
  The number of validators to include in the withdrawal process from the specified starting index.

- **`withdrawAmount`**:
  The amount in gwei to be withdrawn per request. Setting this to `0` triggers a full exit for the validator.

- **`walletPrivkey`**:
  The private key of the wallet initiating the withdrawal requests, necessary for transaction authorization.

- **`withdrawalContract`**:
  The address of the smart contract that handles the withdrawal requests on the blockchain.

- **`txAmount`**, **`txFeeCap`**, **`txTipCap`**, **`txGasLimit`**:
  Transaction parameters including the amount to be transferred, fee cap, tip cap, and gas limit, which dictate the economic aspects of the withdrawal transactions.

- **`clientPattern`**, **`excludeClientPattern`**:
  Regular expressions for selecting or excluding specific client endpoints for sending the withdrawal requests.

- **`awaitReceipt`**:
  Specifies whether to wait for a receipt confirmation for each transaction, confirming execution on the network.

- **`failOnReject`**:
  Determines if the task should fail upon transaction rejection, enhancing error handling and response strategies.

### Outputs

- **`transactionHashes`**:
  A list of hashes for the transactions that have been sent. This output provides identifiers for tracking the transactions on the blockchain.

- **`transactionReceipts`**:
  If `awaitReceipt` is true, this will include the receipts for each transaction, providing details such as success status, gas used, and potentially logs if applicable.

### Defaults

Default settings for the `generate_withdrawal_requests` task:

```yaml
- name: generate_withdrawal_requests
  config:
    limitPerSlot: 0
    limitTotal: 0
    limitPending: 0
    sourcePubkey: ""
    sourceMnemonic: ""
    sourceStartIndex: 0
    sourceStartValidatorIndex: null
    sourceIndexCount: 0
    withdrawAmount: 0
    walletPrivkey: ""
    withdrawalContract: "0x00000961Ef480Eb55e80D19ad83579A64c007002"
    txAmount: "500000000000000000"
    txFeeCap: "100000000000"
    txTipCap: "1000000000"
    txGasLimit: 200000
    clientPattern: ""
    excludeClientPattern: ""
    awaitReceipt: false
    failOnReject: false
```