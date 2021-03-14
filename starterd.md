## Build

```bash
make install
```

## Init chain

```bash
nsd init giansalex --chain-id namechain

nscli config chain-id namechain
nscli config output json
nscli config indent true
nscli config trust-node true

nscli config keyring-backend test 

nscli keys add jack
nscli keys add alice

nsd add-genesis-account $(nscli keys show jack -a) 1000000000uname
nsd add-genesis-account $(nscli keys show alice -a) 1000000000uname
```

Change `stake` default token to `uname` native token at `$HOME/.nsd/config/genesis.json`.     
Also you can change `voting_period` (nanoseconds) in **gov** `genesis.json`

**Params**

> You can change in `genesis.json` or gov proposal.

|Key       | Value  |
|----------|--------| 
|MinPrice  | 100000 |
|BondDenom | `uname`|

Follow:

```bash
nsd gentx --name jack --amount=10000000uname --keyring-backend test
nsd collect-gentxs
```

Run node
```bash
nsd start
```

Optional, run rest server
```bash
nscli rest-server --chain-id nameservice --trust-node
```

## Nameservice commands

```bash
# buy
nscli tx nameservice buy-name jack.id 120000uname --from jack
# change ip resolve
nscli tx nameservice set-name jack.id 8.8.8.8 --from jack
# resolve name
nscli query nameservice resolve jack.id
```


## Proposal

```bash
nscli tx gov submit-proposal param-change ./props/proposal.json --from=jack  --chain-id=namechain

# Get Propsal
nscli query gov proposal 1

# Vote
nscli tx gov vote 1 yes --from=jack  --chain-id=namechain

# List votes
nscli query gov votes 1
# stas votes
nscli query gov tally 1
```