// Copyright 2017 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package chain

// testnetExtraData
/*
UTC--2021-11-15T07-19-41.272626000Z--a45d87cfc1c3ca535d6609746e21e416cb5c1247
private key: 1597a7f6a406b2180e5936797a6d0bc2396896220bd3f448daf58f6056eff1a8
public key: 048ee7932f31a81dd91c9f66775503eb366e411012748bd55d7499c9f9c1cfe8549b75c6275a8183ee6791b610e2a47f6ff489d0f60905ddb5ff2a9071906c7fd1
bls private key: d7d43c0ab866b6ffe37ce4fd37a975b4ce2c68b7df1ac94df81bada453ad6910
bls public key: f9bf370ce720a3d22c72cb3fb13b730604344fb8f4fc171388aed7a2fbbbef26b5e857aa4e553dbe693d7d2caed92b01ca2c9c9d5a10323f270bc1dd6435dca15d58038447f55ea99f830413082792d9cdd90c9e7470e33bca8c1a9babe95280

UTC--2021-11-15T07-19-43.770320000Z--d74d24cc30687cb802c059b5b26e3503fa1feb4f
private key: e05c112f7031b4a86562f3db25f9b5964026bc1f07f5696d7c46eecbeefa34fd
public key: 04a2da8be962ee625efd2572db5ff9546addf9f34795f4c054ff85459b251cf300c44aa58fc539f33e7b4fd4e5f7b429d9edb702f3a16206a7fedcda9b0ce01f54
bls private key: 61cc81f32139bb636c949b27594c096b7b860d32b41d5c6fc0bbc0b8c4da4b08
bls public key: c048189cd5e2cc5695936719f6e40153301adc284e6fc88420e6aee97004bf1cf83421801014c501a76ad397848d5b018899a8727089422af17cb78f77419385c79e3f40736b7caf085f42ced030c1988366a47526801fc108b6d511b6e53e01

UTC--2021-11-15T07-19-45.975328000Z--cf1948fa74fc8c93e77c7d57b4e696f29eb30d3c
private key: 158ac4a6d97672fc4ce3e521cfcdae98b5ca3f6c39275b9d61adde9d9e5de68f
public key: 04c9e1988255827507fa83f09edff814ef5de05498347ce1a2d59d08e5a588b15b44a27fe89d2bc03ecd68cc0d912d761c1bd469ee3c3ffa9e544f79d22ba78c35
bls private key: 6363f0bba77ac948f7154550d325f4c9594e7065edf3abfc7fa095d11c421d09
bls public key: 64aa284431b5bc99418f2c4dc756c768e53b3d26919744daecaab976c5421160afeb938b9c7732b3ad85d7bb8fb367002297a280e1ba8be2952b6c28e08ae304a440fbe61cd8b0b0774b78387599245af01998672ee279dbae2fdfb8cc08ea00

UTC--2021-11-15T07-19-48.067826000Z--b522eab434df6baea887ebd1db4ee9560397515f
private key: e6d966ccf4286a89afbbc1f0d582db3940774600b24e2e57569ad1a3debac0f0
public key: 047885baa9ba610ea1693f3b30a0786101c1c7a4119f781df0af9881c8563d9134ce1a5cd4e7866a9de72ac68e514fcb8408d3f56db19289f960b291b631ca7658
bls private key: 3a23df19aa06efcf377ae35be7dba79b443725a5f020a606d8ba57af3afab404
bls public key: 2c220b971937185b5ba7c7d30e652b76370ff45643740fb894a77e860cd60a48e4290a5605a60fb9867cdb4709ad620035cb191a931146f3f7df29cc0dc47fded20d2139f84030e7bc57dcbd800743a5004b3f5b99e37f5209442684a98d1280
*/

const (
	// extra data in genesis block of main net
	mainnetExtraData = "0x0000000000000000000000000000000000000000000000000000000000000000f90376f854941c0edab88dbb72b119039c4d14b1663525b3ac159416fdbcac4d4cc24dca47b9b80f58155a551ca2af942dc45799000ab08e60b7441c36fcc74060ccbe11946c5938b49bacde73a8db7c3a7da208846898bff5f90208b880136ef6be87de9c925869387782afb4cf19496999c2684709daeb3af8d0b59d800bbe05870789f0f9b3cadababa69f5a00a38bbcba71d99c4c35d671442232c4d3017fd6b99e8356a3e4e985bdfc60bbcb8d939c87976a1ff677d7c42989b379a0b4c0f168a544c892bd2b3ec480e3d6c58c7dddb8d83677ebee2e87ab3660b80b8800a2e37ecad6e69bfec9fec2b345d0f8441a0f63acf8b45c0131a78e5d777d52e0a39404ca85f2c08752c1d4ff8df05c82c7880779d61fe3fabcd4fd682463c0515b1f0217561a6a72bd381da19e34c5560c6eccb08ff83d7d3f4ac6da7f5d1ed15a2780f782c1fa571fa65b99694af559b9df168b1d8745ac3bbc7d3fe550b94b880086fac850f3a9f36e8a5107eab0ba79044043dc2cc6b897cbbd0d4bf805570ff270a98f28e2d2e70b7b2ecc41a4a13e453178354997aa2038852c5945f0564bb02cdf57642881a1b40417fe3620429fc087f8dee6a68e5d7193d3243c38a1f3827d0f4cb616722a1fa78a283a17589d7688a769ade77e9d6417c6e2a9adf59c3b88003fea7bc386ea24aaa19c563a4f26f38cbc2ce172ba2310587405f4f05777fb911a4c3553b7b6529ea02a9da3ae2df6f70c3409105b39e1930d6a6ae8344fc221f5dfb2e73cc8ce434d1af33d95366796bdec26ca7cfcc0a03867fabf471884206db6b9e175a131995bd0c70b93a6f2eec96d831ad0c42d13d334f780d578834f90108b84014d44a97d2fc3ea62b6dcf2bd857079bd261993152f11aef5dd001db68b20d2d1ba45f117b6530a7aec45d7d90fd4e15d2a62f62b706eaa115aa801caeee294bb84015b7bcf0accf839170a5d4621282edcf14f4a438f8e53abcead5f0528cb91cb1135fd4e82ede1493ab1209af122e1dc186c885cc96d2413cbc09a58163b91eb9b8402fd433e93187f6b3d15664ec48073bd73d57c801c4a8bfc1e0e3abd3deefc45619d45ac7ad54df7dda5b8afd6f882c9d9f879dbc6d587f1da5da1751baac729fb8401b037f39d9f8e74b608a898249cc3d156ff1f0051026388366b85a84aac43bb4068275cd909e16b29f1b3bc97e91ec0a8b95a11b8a574cbc2c9ea142d26c8a498080c3808080c3808080"
	// extra data in genesis block of test net
	testnetExtraData = "0x0000000000000000000000000000000000000000000000000000000000000000f90376f854941c0edab88dbb72b119039c4d14b1663525b3ac159416fdbcac4d4cc24dca47b9b80f58155a551ca2af942dc45799000ab08e60b7441c36fcc74060ccbe11946c5938b49bacde73a8db7c3a7da208846898bff5f90208b880136ef6be87de9c925869387782afb4cf19496999c2684709daeb3af8d0b59d800bbe05870789f0f9b3cadababa69f5a00a38bbcba71d99c4c35d671442232c4d3017fd6b99e8356a3e4e985bdfc60bbcb8d939c87976a1ff677d7c42989b379a0b4c0f168a544c892bd2b3ec480e3d6c58c7dddb8d83677ebee2e87ab3660b80b8800a2e37ecad6e69bfec9fec2b345d0f8441a0f63acf8b45c0131a78e5d777d52e0a39404ca85f2c08752c1d4ff8df05c82c7880779d61fe3fabcd4fd682463c0515b1f0217561a6a72bd381da19e34c5560c6eccb08ff83d7d3f4ac6da7f5d1ed15a2780f782c1fa571fa65b99694af559b9df168b1d8745ac3bbc7d3fe550b94b880086fac850f3a9f36e8a5107eab0ba79044043dc2cc6b897cbbd0d4bf805570ff270a98f28e2d2e70b7b2ecc41a4a13e453178354997aa2038852c5945f0564bb02cdf57642881a1b40417fe3620429fc087f8dee6a68e5d7193d3243c38a1f3827d0f4cb616722a1fa78a283a17589d7688a769ade77e9d6417c6e2a9adf59c3b88003fea7bc386ea24aaa19c563a4f26f38cbc2ce172ba2310587405f4f05777fb911a4c3553b7b6529ea02a9da3ae2df6f70c3409105b39e1930d6a6ae8344fc221f5dfb2e73cc8ce434d1af33d95366796bdec26ca7cfcc0a03867fabf471884206db6b9e175a131995bd0c70b93a6f2eec96d831ad0c42d13d334f780d578834f90108b84014d44a97d2fc3ea62b6dcf2bd857079bd261993152f11aef5dd001db68b20d2d1ba45f117b6530a7aec45d7d90fd4e15d2a62f62b706eaa115aa801caeee294bb84015b7bcf0accf839170a5d4621282edcf14f4a438f8e53abcead5f0528cb91cb1135fd4e82ede1493ab1209af122e1dc186c885cc96d2413cbc09a58163b91eb9b8402fd433e93187f6b3d15664ec48073bd73d57c801c4a8bfc1e0e3abd3deefc45619d45ac7ad54df7dda5b8afd6f882c9d9f879dbc6d587f1da5da1751baac729fb8401b037f39d9f8e74b608a898249cc3d156ff1f0051026388366b85a84aac43bb4068275cd909e16b29f1b3bc97e91ec0a8b95a11b8a574cbc2c9ea142d26c8a498080c3808080c3808080"
	devnetExtraData  = "0x0000000000000000000000000000000000000000000000000000000000000000f90376f854941c0edab88dbb72b119039c4d14b1663525b3ac159416fdbcac4d4cc24dca47b9b80f58155a551ca2af942dc45799000ab08e60b7441c36fcc74060ccbe11946c5938b49bacde73a8db7c3a7da208846898bff5f90208b880136ef6be87de9c925869387782afb4cf19496999c2684709daeb3af8d0b59d800bbe05870789f0f9b3cadababa69f5a00a38bbcba71d99c4c35d671442232c4d3017fd6b99e8356a3e4e985bdfc60bbcb8d939c87976a1ff677d7c42989b379a0b4c0f168a544c892bd2b3ec480e3d6c58c7dddb8d83677ebee2e87ab3660b80b8800a2e37ecad6e69bfec9fec2b345d0f8441a0f63acf8b45c0131a78e5d777d52e0a39404ca85f2c08752c1d4ff8df05c82c7880779d61fe3fabcd4fd682463c0515b1f0217561a6a72bd381da19e34c5560c6eccb08ff83d7d3f4ac6da7f5d1ed15a2780f782c1fa571fa65b99694af559b9df168b1d8745ac3bbc7d3fe550b94b880086fac850f3a9f36e8a5107eab0ba79044043dc2cc6b897cbbd0d4bf805570ff270a98f28e2d2e70b7b2ecc41a4a13e453178354997aa2038852c5945f0564bb02cdf57642881a1b40417fe3620429fc087f8dee6a68e5d7193d3243c38a1f3827d0f4cb616722a1fa78a283a17589d7688a769ade77e9d6417c6e2a9adf59c3b88003fea7bc386ea24aaa19c563a4f26f38cbc2ce172ba2310587405f4f05777fb911a4c3553b7b6529ea02a9da3ae2df6f70c3409105b39e1930d6a6ae8344fc221f5dfb2e73cc8ce434d1af33d95366796bdec26ca7cfcc0a03867fabf471884206db6b9e175a131995bd0c70b93a6f2eec96d831ad0c42d13d334f780d578834f90108b84014d44a97d2fc3ea62b6dcf2bd857079bd261993152f11aef5dd001db68b20d2d1ba45f117b6530a7aec45d7d90fd4e15d2a62f62b706eaa115aa801caeee294bb84015b7bcf0accf839170a5d4621282edcf14f4a438f8e53abcead5f0528cb91cb1135fd4e82ede1493ab1209af122e1dc186c885cc96d2413cbc09a58163b91eb9b8402fd433e93187f6b3d15664ec48073bd73d57c801c4a8bfc1e0e3abd3deefc45619d45ac7ad54df7dda5b8afd6f882c9d9f879dbc6d587f1da5da1751baac729fb8401b037f39d9f8e74b608a898249cc3d156ff1f0051026388366b85a84aac43bb4068275cd909e16b29f1b3bc97e91ec0a8b95a11b8a574cbc2c9ea142d26c8a498080c3808080c3808080"
)
