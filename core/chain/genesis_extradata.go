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
	//mainnetExtraData = "0x0000000000000000000000000000000000000000000000000000000000000000f901ebf854941c0edab88dbb72b119039c4d14b1663525b3ac159416fdbcac4d4cc24dca47b9b80f58155a551ca2af942dc45799000ab08e60b7441c36fcc74060ccbe11946c5938b49bacde73a8db7c3a7da208846898bff5f90188b860be77f945929d5dd3fe99aa825df0f5b1e8ea11786333b4492a8624a4d08dcee0e89df327359e8ec3f2d8ae01e938b7003414aa2d6523ffa02fde42b278cbae311fd39f1fbcad8e3188442ea31dee662389599751f8e73b99215cefc2e0003f81b8604f38a71fb13ab20f7bbfc2749ab15d775b7729842d967ca4f4115d1fcb3f378c892d073344f84e2abd8995a16eeee8004f4e588c30261e08a5dae70c581f904ea86b574bfe279222cf6b7913bebb0d3bd6c2bbe2e2ea1d338f145c4d95b99201b8608cf3bfcbfc76e9a99b70cad65ae51f8a8972e3e230445a55c8cf6b96dea7a2d0d970e3545e1316554d5d3b0a53582800ad4de92e3b06b62aa6f7677fdc2885a90b75fd80e2db2775512d8f3d3900aabae5b0525786d65615994b07afe7f69481b8601bbb8eb14a7f5dddc9de3356ce4247dab8e554fa83cd33e663db148b5d2dd14485f090978c84074154b450329de06b018eac04113ede1eedadf891ee862877af92a648c162be62182db90e8c83f8fd154fc14f13676bcb1fe3503260b6261a018080c3808080c3808080"
	mainnetExtraData = "0x0000000000000000000000000000000000000000000000000000000000000000f9047ef854941c0edab88dbb72b119039c4d14b1663525b3ac159416fdbcac4d4cc24dca47b9b80f58155a551ca2af942dc45799000ab08e60b7441c36fcc74060ccbe11946c5938b49bacde73a8db7c3a7da208846898bff5f9020cb8810185937aebca906731bbd7dddf9e8439f27d4e187e67e8d4eb78db20679fef5d8a776a02e1da3e676afa7ccb4650bfc8759a223d4f0b7c9afc9a2a55e97a36777155854485d24ea91dd4f06cf0c447e1b32379fba24acd34f24a3a292429a582de677158128ffbb1b761379adee7fc0a16dc3f026540b0632a11cb75c382be1eb9b8810188ea7e8925a6aba026ee30141627577c38831b5d6c928e15629668c41c821b8a3bd71b4317c4eed85f9735b64d7d2696361738b7499e0cc7f9a54b62d76fd1370a250417ad7491ff4f368f915a05f5997c6fd7fdc20788bda2a21c7180845ef380dbaac5115541f5569dca7c52e9abe10770292267da5bbc5de12a98c243d679b881018e6d48f3f550e72d5cf577458856d2710870d2b9ccff500d04894fe234cf69ec8039585715c879d22001545e696dd8f1a6a90a1e8aa6e256c1b3d124c0286d0b372ba11433d71c84d3176eea1d824e8118f70c5ba6b533909ea9503bd14343491b5d4226d57ac3b62142f9c8c2adc1266e91bb6ac996fac8d85a2d41bd1212c1b881010a4ce0564f11ccb6051cb88118a0f86ee09c95cf17746fec303088633a5360d5643b8fe345f62c1d25d4a8327507cbd6d0cee2daf405120ece3ce164eec7446880ae5f8fa56ee8b2f483dd6a1924d855ccc0197d58972f545a2cab809c3507738f26d1c1bdcd87888176e0ea805563a613da12cc8b4b21e481c30f62fbd9a4bff9020cb8810185937aebca906731bbd7dddf9e8439f27d4e187e67e8d4eb78db20679fef5d8a776a02e1da3e676afa7ccb4650bfc8759a223d4f0b7c9afc9a2a55e97a36777155854485d24ea91dd4f06cf0c447e1b32379fba24acd34f24a3a292429a582de677158128ffbb1b761379adee7fc0a16dc3f026540b0632a11cb75c382be1eb9b8810188ea7e8925a6aba026ee30141627577c38831b5d6c928e15629668c41c821b8a3bd71b4317c4eed85f9735b64d7d2696361738b7499e0cc7f9a54b62d76fd1370a250417ad7491ff4f368f915a05f5997c6fd7fdc20788bda2a21c7180845ef380dbaac5115541f5569dca7c52e9abe10770292267da5bbc5de12a98c243d679b881018e6d48f3f550e72d5cf577458856d2710870d2b9ccff500d04894fe234cf69ec8039585715c879d22001545e696dd8f1a6a90a1e8aa6e256c1b3d124c0286d0b372ba11433d71c84d3176eea1d824e8118f70c5ba6b533909ea9503bd14343491b5d4226d57ac3b62142f9c8c2adc1266e91bb6ac996fac8d85a2d41bd1212c1b881010a4ce0564f11ccb6051cb88118a0f86ee09c95cf17746fec303088633a5360d5643b8fe345f62c1d25d4a8327507cbd6d0cee2daf405120ece3ce164eec7446880ae5f8fa56ee8b2f483dd6a1924d855ccc0197d58972f545a2cab809c3507738f26d1c1bdcd87888176e0ea805563a613da12cc8b4b21e481c30f62fbd9a4bf8080c3808080c3808080"
	// extra data in genesis block of test net
	testnetExtraData = "0x0000000000000000000000000000000000000000000000000000000000000000f9047ef854941c0edab88dbb72b119039c4d14b1663525b3ac159416fdbcac4d4cc24dca47b9b80f58155a551ca2af942dc45799000ab08e60b7441c36fcc74060ccbe11946c5938b49bacde73a8db7c3a7da208846898bff5f9020cb8810185937aebca906731bbd7dddf9e8439f27d4e187e67e8d4eb78db20679fef5d8a776a02e1da3e676afa7ccb4650bfc8759a223d4f0b7c9afc9a2a55e97a36777155854485d24ea91dd4f06cf0c447e1b32379fba24acd34f24a3a292429a582de677158128ffbb1b761379adee7fc0a16dc3f026540b0632a11cb75c382be1eb9b8810188ea7e8925a6aba026ee30141627577c38831b5d6c928e15629668c41c821b8a3bd71b4317c4eed85f9735b64d7d2696361738b7499e0cc7f9a54b62d76fd1370a250417ad7491ff4f368f915a05f5997c6fd7fdc20788bda2a21c7180845ef380dbaac5115541f5569dca7c52e9abe10770292267da5bbc5de12a98c243d679b881018e6d48f3f550e72d5cf577458856d2710870d2b9ccff500d04894fe234cf69ec8039585715c879d22001545e696dd8f1a6a90a1e8aa6e256c1b3d124c0286d0b372ba11433d71c84d3176eea1d824e8118f70c5ba6b533909ea9503bd14343491b5d4226d57ac3b62142f9c8c2adc1266e91bb6ac996fac8d85a2d41bd1212c1b881010a4ce0564f11ccb6051cb88118a0f86ee09c95cf17746fec303088633a5360d5643b8fe345f62c1d25d4a8327507cbd6d0cee2daf405120ece3ce164eec7446880ae5f8fa56ee8b2f483dd6a1924d855ccc0197d58972f545a2cab809c3507738f26d1c1bdcd87888176e0ea805563a613da12cc8b4b21e481c30f62fbd9a4bff9020cb8810185937aebca906731bbd7dddf9e8439f27d4e187e67e8d4eb78db20679fef5d8a776a02e1da3e676afa7ccb4650bfc8759a223d4f0b7c9afc9a2a55e97a36777155854485d24ea91dd4f06cf0c447e1b32379fba24acd34f24a3a292429a582de677158128ffbb1b761379adee7fc0a16dc3f026540b0632a11cb75c382be1eb9b8810188ea7e8925a6aba026ee30141627577c38831b5d6c928e15629668c41c821b8a3bd71b4317c4eed85f9735b64d7d2696361738b7499e0cc7f9a54b62d76fd1370a250417ad7491ff4f368f915a05f5997c6fd7fdc20788bda2a21c7180845ef380dbaac5115541f5569dca7c52e9abe10770292267da5bbc5de12a98c243d679b881018e6d48f3f550e72d5cf577458856d2710870d2b9ccff500d04894fe234cf69ec8039585715c879d22001545e696dd8f1a6a90a1e8aa6e256c1b3d124c0286d0b372ba11433d71c84d3176eea1d824e8118f70c5ba6b533909ea9503bd14343491b5d4226d57ac3b62142f9c8c2adc1266e91bb6ac996fac8d85a2d41bd1212c1b881010a4ce0564f11ccb6051cb88118a0f86ee09c95cf17746fec303088633a5360d5643b8fe345f62c1d25d4a8327507cbd6d0cee2daf405120ece3ce164eec7446880ae5f8fa56ee8b2f483dd6a1924d855ccc0197d58972f545a2cab809c3507738f26d1c1bdcd87888176e0ea805563a613da12cc8b4b21e481c30f62fbd9a4bf8080c3808080c3808080"
	devnetExtraData  = "0x0000000000000000000000000000000000000000000000000000000000000000f9047ef854941c0edab88dbb72b119039c4d14b1663525b3ac159416fdbcac4d4cc24dca47b9b80f58155a551ca2af942dc45799000ab08e60b7441c36fcc74060ccbe11946c5938b49bacde73a8db7c3a7da208846898bff5f9020cb8810185937aebca906731bbd7dddf9e8439f27d4e187e67e8d4eb78db20679fef5d8a776a02e1da3e676afa7ccb4650bfc8759a223d4f0b7c9afc9a2a55e97a36777155854485d24ea91dd4f06cf0c447e1b32379fba24acd34f24a3a292429a582de677158128ffbb1b761379adee7fc0a16dc3f026540b0632a11cb75c382be1eb9b8810188ea7e8925a6aba026ee30141627577c38831b5d6c928e15629668c41c821b8a3bd71b4317c4eed85f9735b64d7d2696361738b7499e0cc7f9a54b62d76fd1370a250417ad7491ff4f368f915a05f5997c6fd7fdc20788bda2a21c7180845ef380dbaac5115541f5569dca7c52e9abe10770292267da5bbc5de12a98c243d679b881018e6d48f3f550e72d5cf577458856d2710870d2b9ccff500d04894fe234cf69ec8039585715c879d22001545e696dd8f1a6a90a1e8aa6e256c1b3d124c0286d0b372ba11433d71c84d3176eea1d824e8118f70c5ba6b533909ea9503bd14343491b5d4226d57ac3b62142f9c8c2adc1266e91bb6ac996fac8d85a2d41bd1212c1b881010a4ce0564f11ccb6051cb88118a0f86ee09c95cf17746fec303088633a5360d5643b8fe345f62c1d25d4a8327507cbd6d0cee2daf405120ece3ce164eec7446880ae5f8fa56ee8b2f483dd6a1924d855ccc0197d58972f545a2cab809c3507738f26d1c1bdcd87888176e0ea805563a613da12cc8b4b21e481c30f62fbd9a4bff9020cb8810185937aebca906731bbd7dddf9e8439f27d4e187e67e8d4eb78db20679fef5d8a776a02e1da3e676afa7ccb4650bfc8759a223d4f0b7c9afc9a2a55e97a36777155854485d24ea91dd4f06cf0c447e1b32379fba24acd34f24a3a292429a582de677158128ffbb1b761379adee7fc0a16dc3f026540b0632a11cb75c382be1eb9b8810188ea7e8925a6aba026ee30141627577c38831b5d6c928e15629668c41c821b8a3bd71b4317c4eed85f9735b64d7d2696361738b7499e0cc7f9a54b62d76fd1370a250417ad7491ff4f368f915a05f5997c6fd7fdc20788bda2a21c7180845ef380dbaac5115541f5569dca7c52e9abe10770292267da5bbc5de12a98c243d679b881018e6d48f3f550e72d5cf577458856d2710870d2b9ccff500d04894fe234cf69ec8039585715c879d22001545e696dd8f1a6a90a1e8aa6e256c1b3d124c0286d0b372ba11433d71c84d3176eea1d824e8118f70c5ba6b533909ea9503bd14343491b5d4226d57ac3b62142f9c8c2adc1266e91bb6ac996fac8d85a2d41bd1212c1b881010a4ce0564f11ccb6051cb88118a0f86ee09c95cf17746fec303088633a5360d5643b8fe345f62c1d25d4a8327507cbd6d0cee2daf405120ece3ce164eec7446880ae5f8fa56ee8b2f483dd6a1924d855ccc0197d58972f545a2cab809c3507738f26d1c1bdcd87888176e0ea805563a613da12cc8b4b21e481c30f62fbd9a4bf8080c3808080c3808080"
)
