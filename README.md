# Gemer: A Verifiable Stateless Environment Builder
### go-gemer - Proof of Concept implementation of [Gemer](https://gemer.io) written in Go

**WARNING** This implementation of the [Gemer](https://gemer.io) paper is strictly meant to be a simple proof-of-concept. I am planning to reimplement this from scratch, if I ever decide to take this to production level.

### What is [Gemer](https://gemer.io)?

**Gemer** is a framework that builds a *verifiable stateless computing environment*. In other words, this allows developers to build fully decentralized applications that truly respect user's privacy with little to no performance impact.

**Think about it this way:** it's like BitTorrent, but instead of just downloading static files you can run the apps you love on the network. Also like BitTorrent, it can run apps without a centralized server or a cloud. Users can keep their data on their own devices or encrypted storage, without leaking that to the service provider. Even better, if your app's backend is completely WebAssembly-compatible you may be able to port over the entire app on it - with no to minimal impact on performance or user experience. And you get full control over your entire codebase, so update and make changes to your app just like you are used to - no dealing with immutable chunks of code that should run forever on a blockchain (and not being affected by security vulnerabilities that cannot be fixed, unless you design it to be upgradable in the first place).

**Even simpler:** it's BitTorrent + AWS, or BitTorrent + WebAssembly - depends on whichever way you choose to use or see it.

This is architecturally different with blockchains, distributed ledgers or cryptocurrencies like Bitcoin or Ethereum - it's something completely new. For a detailed explanation, please refer to the [website](https://gemer.io) or the [paper](https://papers.unifiedh.com/Gemer/gemer.pdf).

### Contributing

If you are interested in this concept and would like to contribute, please feel free to leave an issue or a pull request. I will not put any restrictions and/or rules for contributing to this project, as this is currently not a production-level solution.

### License

This project is licensed under the MIT License.
