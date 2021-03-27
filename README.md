[<img src="https://avatars2.githubusercontent.com/u/24763891?s=400&u=c1150e7da5667f47159d433d8e49dad99a364f5f&v=4"  width="256px" height="256px" align="right" alt="Multiverse OS Logo">](https://github.com/multiverse-os)

## Multiverse: Wormholes Router, easily construct overlay network bootstrapped P2P/F2F DHT(s)
**URL** [multiverse-os.org](https://multiverse-os.org)

### Introduction
A bare-bones skeleton-like router implemented primarily as a library, intended to function as a drop in P2P networking support. 

#### Software-defined Tunnels
The goal is to enable developers create simple expressive scripts that a ruleset for specific or non-specific tunnel creation. This functionality combines the functionality of Tor and I2C, and extends it further into the limits of the imagination of interested developers. The creation of bots that use the network in erratic ways to obstuficate actual traffic, or missing traffic indicating you are not home, or anonymize your traffic, obstuficating your DNS lookups by sending fake ones to your ISP while doing all real lookups over encrypted connections tunneled through other protocols, sending data through the distance of time between each different DNS request made to a server at the end of a tunnel. 

#### Foundation for multiparty computation
Importantly, the creation of single hop to create rendevous relay node powered direct connection with peers allows for access to submit third party receipts to the network to track contributions.

This allows for third parties to verify data provided from a peer, but it also provides a way to submit third party verified data to the network, and have it be written. 

To scaffold P2P (peer-to-peer, or public) networks, and private F2F (friend-to-friend) network, several software defined tunnels need to be connected into a network. And the selection of which tunnel to use, and its specific ruleset, will influence how development of the network continues. 

### Creating P2P Networks from Tunnels
Using protocol definitions to provide ways to pay the network, by asking the network the network, if you wanted to contribute in this way (for example, mailing cash, paypal, etc), it wouensuring incentives are divided up in a way to create incentivization in an automated, and transparent way which prevents interception of money sent to the network, while eliminating any need for administration. 

This will enable developers to drop in networking logic that can provide a solution to a wide variety of use cases, from direct redevous based direct communication capable of bypassing NAT restrictions; and all of this is defined used a library that provides both the Go library, but also a Ruby library that interacts with the Go process and builds binary objects that the Go project can build structs from. 

Once friends can direct connect together, this overlay network can function to provide the wider wormholes overlay network with a map of hops that are possibly faster and more secure. This information along with a ping map, or a geolocal graph representation based on latency between nodes instead of simply distance apart, since these two values may appear at first to be related or the same they are often confusing. The ping map would done to known friends, and trusted peers, always keeping any location data in encrypted GeoJSON. The distance in latency and location can be measured and used to construct more secure tunnels, and using anonymized friends of friends list, focusing on finding the fastest and most secure ways to enter and exit the overlay network are prioritized and relay functionality taking away much of their power using multiparty computation to restrict the knowledge of their activity.


**Please stop that is a lot of text**\n
tl/dr: *Using this library a developer can define a set of rules, for the creation of the entrance node, all relay nodes, and an exit node if the internet is the final destination utilizing an onion routing overlay network inspired but independent from the first generation of onion networks.* 



### Project Layout
The project is implemented primarily as a library, then each of the clients for each device can be built independently, or can be later combined using a broadly supported framework such as QT. 


```
.
├── cmd
│   ├── wormholed
│   ├── wormholes-android
│   ├── wormholes-cli
│   ├── wormholes-ios
│   ├── wormholes-osx
│   ├── wormholes-webui
│   └── wormholes-windows
├── data
│   ├── data.go
│   └── mongo
├── protocol
│   ├── crypto
│   ├── network
│   ├── os
│   └── protocol.go
├── README.md
├── router
│   ├── dhcp
│   ├── ip
│   ├── overlay
│   ├── p2p
│   ├── router.go
│   └── tables
├── services
│   ├── chat
│   ├── http
│   ├── services.go
│   └── users
└── wormholes.go
```


### Client Development
*For now the CLI client will take priority during development, focus on a webui will be prioritized so that a UI can be implemented using HTML/CSS (explicitly without Javascript) to provide content and style it within a window. And an IOS likely native client in Go is an early test candidate*.

Implementing even a minimal javascript to provide DOM manipulation seems unlikely with how easy it is to manipulate a DOM from Go. If construction of the HTML/CSS elements could be done in parallel in a engine like Pixels, or directly written to the framebuffer, that would be the most ideal solution for a broadly supported UI that works essentially on every device. However until then, this organization includes code for both a Firefox, Chromium, and Webkit. And there is the possibility of including the binary for the device, and specifically use an extremely minimal version of the web engine so that it has less ability to be a vector, while running much faster. 


#### Specification Clients
Two specification clients will be built, one providing an example for the most minimal feature set possible that still properly and generously interacts with the network. The minimal version, that will likely provide the foundation to a secure server client focusing on minimal design while providing a complete feature set, or near complete. 

The Go client will be the primary general-use client, taking full advantage of the nested P2P networks, the data stored within it, and other overlay networks to be a framework for a variety of secure decentralized services that function to provide service to a small group to massive social network. 

The CLI client will provide direct access to the OHT databases

#### CLI Client
This CLI and console client will implement the full featureset, so that if another application wanted to interact with it, or an administrator, test networks, private networks, and public networks are all available: 

```
**Console Client**
wormholes> /help
COMMANDS:
  CONFIG:
    /config                      - List configuration values
    /set [config] [option]       - Change configuration options
    /unset [config]              - Unset configuration option
    /save                        - Save configuration values

  TOR:
    /tor [start|stop]            - Start or stop Tor process
    /newtor                      - Obtain new Tor identity 
    /newonions                   - Obtain new onion address
    // TODO: Needs to add support for I2C and wormholes own overlay network
    // TODO: Experiment with securing writing to several interconnected DHTs by using ring signatures. Enabling sharded data or different tables of data to be securely modified in a way that is resistant to race conditions

  NETWORK:
    /peers                       - List all connected peers
    /successor                   - Next peer in identifier ring (Not Implemented)
    /predecessor                 - Previous peer in identifier ring (Not Implemented)
    /ftable                      - List ftable peers (Not Implemented)
    /create                      - Create new ring (Not Implemented)
    /connect [onion address]     - Join ring containing peer with [onion address]
    /lookup [id]                 - Find onion address of account with [id] (Not Implemented)
    /ping [onion address]        - Ping peer (Not Implemented)
    /ringcast [message]          - Message every peer in ring (Not Implemented)

  DHT:
    /put [key] [value]           - Put key and value into database (Not Implemented)
    /get [key]                   - Get value of key (Not Implemented)
    /delete [key]                - Delete key and its value from database (Not Implemented)

  WEBUI:
    /webui [start|stop]          - Start or stop webUI server

  ACCOUNT:
    /accounts                    - List all local accounts (Not Implemented)
    /generate                    - Generate new account key pair (Not Implemented)
    /delete                      - Delete an account key pair (Not Implemented)
    /sign [id] [message]         - Sign with account key pair (Not Implemented)
    /verify [id] [message]       - Verify a signed message with key pair (Not Implemented)
    /encrypt [id] [message]      - Encrypt a message with key pair (Not Implemented)
    /decrypt [id] [message]      - Decrypt a message with key pair (Not Implemented)

  CONTACTS:
    /contacts                    - List all saved contacts (Not Implemented)
    /request [id] [message]      - Send [message] requesting account with [id] to add your id to their contacts (Not Implemented)
    /add [id]                    - Add account to contacts (Not Implemented)
    /rm [id]                     - Remove account from contacts (Not Implemented)
    /whisper [id] [message]      - Direct message peer (Not Implemented)
    /contactcast [message]       - Message all contacts (Not Implemented)

  CHANNELS:
    /channels                    - List all known channels (Not Implemented)
    /channel                     - Generates a new channel (Not Implemented)
    /join [id]                   - Join channel with id (Not Implemented)
    /leave [id]                  - Leave channel with id (Not Implemented)
    /channelcast [id] [message]  - Message all channel subscribers (Not Implemented)

    /quit                        - Quit oht console
```

