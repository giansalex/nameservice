## Build

```bash
make install
```

## Init chain

```bash
nsd init giansalex --chain-id namechain

nsd keys add jack
nsd keys add alice

nsd add-genesis-account $(nsd keys show jack -a) 1000000000uname
nsd add-genesis-account $(nsd keys show alice -a) 1000000000uname
```

**Params**

> You can change in `genesis.json` or gov proposal.

|Key       | Value  |
|----------|--------| 
|MinPrice  | 10     |
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

## Nameservice commands

```bash
# buy
nsd tx nameservice buy-name jack.id 12uname --from jack
# change ip resolve
nsd tx nameservice set-name jack.id 8.8.8.8 --from jack
# resolve name
nsd query nameservice resolve jack.id
```


## Proposal

```bash
nsd tx gov submit-proposal param-change ./props/proposal.json --from=jack  --chain-id=namechain

# Get Propsal
nsd query gov proposal 1

# Vote
nsd tx gov vote 1 yes --from=jack

# List votes
nsd query gov votes 1
# stas votes
nsd query gov tally 1
```