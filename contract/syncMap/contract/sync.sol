// SPDX-License-Identifier: Apache-2.0
pragma solidity >=0.7.1;

import "../lib/RLPReader.sol";
import "../lib/RLPEncode.sol";
import "../bn256/BlsSignatureTest.sol";

contract sync {

    struct blockHeader{
        //bytes32 parentHash;
        address coinbase;
        //bytes32 root;
        //bytes32 txHash;
        //bytes32 receipHash;
        //bytes bloom;
        uint256 number;
        //uint256 gasLimit;
        //uint256 gasUsed;
        //uint256 time;
        bytes extraData;
        //bytes32 mixDigest;
        //bytes nonce;
        //uint256 baseFee;
    }
    mapping(uint256 => bytes) private allHeader;

    struct istanbulAggregatedSeal{
        uint256   round;
        bytes     signature;
        uint256   bitmap;
    }

    struct istanbulExtra{
        //address[] validators;
        bytes  seal;
        istanbulAggregatedSeal  aggregatedSeal;
        istanbulAggregatedSeal  parentAggregatedSeal;
        uint256  removeList;
        bytes[]  addedPubKey;
    }

    mapping(uint256 => bytes[]) private allkey;
    uint256 nowEpoch;
    uint256 nowNumber;
    address rootAccount;
    uint256 epochLength;
    uint256 maxSyncNum;
    uint256 keyNum;

    using RLPReader for bytes;
    using RLPReader for uint;
    using RLPReader for RLPReader.RLPItem;
    using RLPReader for RLPReader.Iterator;

    event setParams(string s,uint256 v);
    event setParams(string s,bytes v);
    event log(string s,bool e);

    constructor(bytes memory firstBlock,uint _epochLength) {
        rootAccount = msg.sender;
        epochLength = _epochLength;
        maxSyncNum = 10;
        nowEpoch = 0;
        initFirstBlock(firstBlock);
    }

    //the first block is used for init params,
    //it was operated specially.
    function initFirstBlock(bytes memory firstBlock) private{
        blockHeader memory bh = decodeHeaderPart1(firstBlock);
        bytes memory extra = splitExtra(bh.extraData);
        istanbulExtra memory ist = decodeExtraData(extra);

        keyNum = ist.addedPubKey.length;
        nowNumber = bh.number;
        require(nowNumber == 0);
        allkey[nowEpoch] = new bytes[](keyNum);

        for(uint8 i = 0;i<keyNum;i++){
            allkey[nowEpoch][i] = ist.addedPubKey[i];
        }
        allHeader[nowNumber] = firstBlock;
    }

    //
    function setBLSPublickKeys(bytes[] memory keys,uint256 epoch) public {
        require(msg.sender == rootAccount, "onlyRoot");
        emit setParams("current epoch",epoch);
        allkey[epoch] = new bytes[](keys.length);
        for (uint i=0;i<keys.length;i++){
            emit setParams("setBLSPublickKey",keys[i]);
            allkey[epoch][i] = keys[i];
        }
    }

    function setMaxSyncNum(uint8 max) public{
        require(msg.sender == rootAccount, "onlyRoot");
        emit setParams("setMaxSyncNum",max);
        maxSyncNum = max;
    }

    function checkNowParams() public view returns(uint,uint,uint,uint){
        //require(msg.sender == rootAccount, "onlyRoot");
        return (maxSyncNum,nowEpoch,nowNumber,keyNum);
    }

    function checkBLSPublickKeys(uint256 epoch) public view returns(bytes[] memory){
        //require(msg.sender == rootAccount, "onlyRoot");
        return allkey[epoch];
    }

    function checkBlockHeader(uint256 number) public view returns(bytes memory){
        //require(msg.sender == rootAccount, "onlyRoot");
        return allHeader[number];
    }

    //function checkExtraData(uint256 number) public view returns(istanbulExtra memory){
    //    require(msg.sender == rootAccount, "onlyRoot");
    //    return allExtra[number];
    //}

    function decodeHeaderPart1(bytes memory rlpBytes)public pure returns(blockHeader memory bh){
        RLPReader.RLPItem[] memory ls = rlpBytes.toRlpItem().toList();
        //RLPReader.RLPItem memory item0 = ls[0]; //parentBlockHash
        RLPReader.RLPItem memory item1 = ls[1]; //coinbase
        //RLPReader.RLPItem memory item2 = ls[2]; //root
        //RLPReader.RLPItem memory item3 = ls[3]; //txHash
        //RLPReader.RLPItem memory item4 = ls[4]; //receipHash
        RLPReader.RLPItem memory item6 = ls[6]; //number
        RLPReader.RLPItem memory item10 = ls[10]; //extra

        //bh.parentHash = bytes32(item0.toBytes());
        bh.coinbase = item1.toAddress();
        //bh.root = bytes32(item2.toBytes());
        //bh.txHash = bytes32(item3.toBytes());
        //bh.receipHash = bytes32(item4.toBytes());
        bh.number = item6.toUint();
        bh.extraData = item10.toBytes();
        return bh;
    }

    //function decodeHeaderPart2(bytes memory rlpBytes,blockHeader memory bh)public pure returns(blockHeader memory){
    //RLPReader.RLPItem[] memory ls = rlpBytes.toRlpItem().toList();
    //RLPReader.RLPItem memory item5 = ls[5]; //bloom
    //RLPReader.RLPItem memory item7 = ls[7]; //gasLimit
    //RLPReader.RLPItem memory item8 = ls[8]; //gasUsed
    //RLPReader.RLPItem memory item9 = ls[9]; //time
    //RLPReader.RLPItem memory item11 = ls[11]; //mixDigest
    //RLPReader.RLPItem memory item12 = ls[12]; //nonce
    //RLPReader.RLPItem memory item13 = ls[13]; //baseFee

    //bh.bloom = item5.toBytes();
    //bh.gasLimit = item7.toUint();
    //bh.gasUsed = item8.toUint();
    //bh.time = item9.toUint();
    //bh.mixDigest  = bytes32(item11.toBytes());
    //bh.nonce  = item12.toBytes();
    //bh.baseFee = item13.toUint();
    //return bh;
    //}


    function verifymoreHeaders(bytes[] memory moreRlpHeader)public returns(uint,bool){
        require(maxSyncNum > moreRlpHeader.length);
        for(uint i=0;i<moreRlpHeader.length;i++){
            bool ret = verifyHeader(moreRlpHeader[i]);
            if (ret == false){
                return (i,false);
            }
        }
        return (moreRlpHeader.length,true);
    }

    //the input data is the header after rlp encode within seal by proposer and aggregated seal by validators.
    function verifyHeader(bytes memory rlpHeader) public returns(bool){
        bool ret = true;

        //it only decode data about validation,so that reduce storage and calculation.
        blockHeader memory bh = decodeHeaderPart1(rlpHeader);
        //bh = decodeHeaderPart2(rlpHeader,bh);
        bytes memory extra = splitExtra(bh.extraData);
        istanbulExtra memory ist = decodeExtraData(extra);

        //the escaMsg is the hash of the header without seal by proposer and aggregated seal by validators.
        bytes memory blsMsg = cutSeal(rlpHeader,ist.seal);
        bytes memory rlpAggregatedSeal = encodeAgg(ist.aggregatedSeal.signature,ist.aggregatedSeal.round,ist.aggregatedSeal.bitmap);
        bytes memory ecdsaMsg = cutSeal(blsMsg,rlpAggregatedSeal);
        bytes32 ecdsaSignMsg = keccak256(abi.encodePacked(ecdsaMsg));
        //the esca seal signed by proposer
        ret = verifySign(ist.seal,ecdsaSignMsg,bh.coinbase);
        if (ret == false) {
            revert("verifyEscaSign fail");
            //return false;
        }
        // emit log("verifyEscaSign pass",true);


        // //the blockHash is the hash of the header without aggregated seal by validators.
        // //bytes memory blockHash = cutAgg(rlpHeader);
        // if (bh.number%epochLength == 0){
        //     bytes32 blsSignMsg = keccak256(abi.encodePacked(blsMsg));
        //     //ret = verifyAggregatedSeal(allkey[nowEpoch],ist.aggregatedSeals.Signature,blsSignMsg);
        //     //it need to update validators at first block of new epoch.
        //     changeValidators(ist.removeList,ist.addedPubKey);
        //     emit log("changeValidators pass",true);
        // }else{
        //     //ret = verifyAggregatedSeal(allkey[nowEpoch],ist.aggregatedSeals.Signature,blsSignMsg);
        // }
        // // if (ret == false) {
        // //     revert("verifyBlsSign fail");
        // //     //return false;
        // // }

        // //the parent seal need to pks of last epoch to verify parent seal,if block number is the first block or the second block at new epoch.
        // //because, the parent seal of the first block and the second block is signed by validitors of last epoch.
        // //and it need to not verify, when the block number is less than 2, the block is no parent seal.
        // if (bh.number > 1) {
        //     if ((bh.number-1)%epochLength == 0 || (bh.number)%epochLength == 0){
        //         //ret = verifyAggregatedSeal(allkey[nowEpoch-1],ist.parentAggregatedSeals.Signature,bh.parentHash);
        //     }else{
        //         //ret = verifyAggregatedSeal(allkey[nowEpoch],ist.parentAggregatedSeals.Signature,bh.parentHash);
        //     }
        //     if (ret == false) {
        //         revert("verifyBlsSign fail");
        //         //return false;
        //     }
        // }

        // nowNumber = nowNumber + 1;
        // //if(nowNumber+1 != bh.number){
        // //    revert("number error");
        // //    //return false;
        // //}
        // allHeader[nowNumber] = rlpHeader;
        // emit log("verifyHeader pass",true);
        return ret;
    }

    function decodeExtraData(bytes memory rlpBytes) public pure returns(istanbulExtra memory ist){
        RLPReader.RLPItem[] memory ls = rlpBytes.toRlpItem().toList();
        RLPReader.RLPItem memory item1 = ls[1];
        RLPReader.RLPItem memory item2 = ls[2];
        RLPReader.RLPItem memory item3 = ls[3];
        RLPReader.RLPItem memory item4 = ls[4];
        RLPReader.RLPItem memory item5 = ls[5];

        //Usually, the length of BLS pk is 98 bytes.
        //According to its length, it can calculate the number of pk.
        if (item1.len > 98){
            uint num = (item1.len - 2)/98;
            ist.addedPubKey = new bytes[](num);
            for(uint i=0;i<num;i++){
                ist.addedPubKey[i] = item1.toList()[i].toBytes();
            }
        }

        ist.removeList = item2.toUint();
        ist.seal = item3.toBytes();
        ist.aggregatedSeal.round = item4.toList()[2].toUint();
        ist.aggregatedSeal.signature = item4.toList()[1].toBytes();
        ist.aggregatedSeal.bitmap = item4.toList()[0].toUint();
        ist.parentAggregatedSeal.round = item5.toList()[2].toUint();
        ist.parentAggregatedSeal.signature = item5.toList()[1].toBytes();
        ist.parentAggregatedSeal.bitmap = item5.toList()[0].toUint();

        return  ist;
    }

    //the function will select legal validators from old validator and add new validators.
    function changeValidators(uint256 removedVal,bytes[] memory addVal) public view returns(bytes[] memory ret){
        (uint[] memory list,uint8 oldVal) = readRemoveList(removedVal);
        ret = new bytes[](oldVal+addVal.length);
        uint j=0;
        //if value is 1, the related address will be not validaor at nest epoch.
        for(uint i=0;i<list.length;i++){
            if (list[i] == 0){
                ret[j] = allkey[nowEpoch][i];
                j = j + 1;
            }
        }
        for(uint i=0;i<addVal.length;i++){
            ret[j] = addVal[i];
            j = j + 1;
        }
        //require(j<101,"the number of validators is more than 100")
        return ret;
    }

    //it return binary data and the number of validator in the list.
    function readRemoveList(uint256 r) public view returns(uint[] memory ret,uint8 sum){
        //the function transfer uint to binary.
        sum = 0;
        ret = new uint[](keyNum);
        for(uint i=0;r>0;i++){
            if (r%2 == 1){
                r = (r-1)/2;
                ret[i] = 1;
            }else{
                r = r/2;
                ret[i] = 0;
                sum = sum + 1;
            }
        }
        //the current array is inverted.it needs to count down.
        for(uint i=0;i<ret.length/2;i++) {
            uint temp = ret[i];
            ret[i] = ret[ret.length-1-i];
            ret[ret.length-1-i] = temp;
        }
        return (ret,sum);
    }

    function verifySign(bytes memory seal,bytes32 hash,address coinbase) public pure returns (bool){
        //Signature storaged in extraData sub 27 after proposer signed.
        //So signature need to add 27 when verify it.
        (bytes32 r, bytes32 s, uint8 v) = splitSignature(seal);
        v=v+27;
        return coinbase == ecrecover(hash, v, r, s);
    }

    function splitSignature(bytes memory sig) internal pure returns (bytes32 r,bytes32 s,uint8 v){
        require(sig.length == 65, "invalid signature length");
        assembly {
            r := mload(add(sig, 32))
            s := mload(add(sig, 64))
            v := byte(0, mload(add(sig, 96)))
        }
    }

    function splitExtra(bytes memory extra) internal pure returns (bytes memory newExtra){
        //extraData rlpcode is storaged from No.32 byte to latest byte.
        //So, the extraData need to reduce 32 bytes at the beginning.
        newExtra = new bytes(extra.length - 32);
        uint n = 0;
        for(uint i=32;i<extra.length;i++){
            newExtra[n] = extra[i];
            n = n + 1;
        }
        return newExtra;
    }

    function encodeAgg(bytes memory signature,uint round,uint bitmap) public pure returns (bytes memory output){
        bytes memory output1 = RLPEncode.encodeUint(round);//round
        bytes memory output2 = RLPEncode.encodeBytes(signature);//signature
        bytes memory output3 = RLPEncode.encodeUint(bitmap);//bitmap

        uint index = 0;
        //

        output = new bytes(output1.length+output2.length+output3.length);
        for(uint i=0;i<output1.length;i++){
            output[index] = output1[i];
            index++;
        }
        for(uint i=0;i<output2.length;i++){
            output[index] = output2[i];
            index++;
        }
        for(uint i=0;i<output3.length;i++){
            output[index] = output3[i];
            index++;
        }
    }

    function setPre(uint hbLen,uint sealLen)public pure returns(bytes memory pre,uint len){
        if(sealLen<257 && hbLen>257){
            len = hbLen-sealLen-2;
            pre = new bytes(2);
            pre[0] = bytes1(uint8(248));
            pre[1] = bytes1(uint8(len)); //h-s+nil(1)-pre(3)
        }else if(sealLen>257 && hbLen>257){
            len = hbLen-sealLen-2;
            pre = new bytes(3);
            pre[0] = bytes1(uint8(249));
            pre[1] = bytes2(uint16(len))[0]; //h-s+nil(1)-pre(3)
            pre[2] = bytes2(uint16(len))[1];
        }else{
            len = hbLen-sealLen-1;
            pre = new bytes(2);
            pre[0] = bytes1(uint8(248));
            pre[1] = bytes1(uint8(len)); //h-s+nil(1)-pre(2)
        }
        return (pre,len+pre.length);//dataLen+preLen
    }

    function bfsearch(bytes memory hb,bytes memory seal)public pure returns(uint){
        for (uint i=0;i<hb.length;i++) {
            for (uint j=0;j<seal.length;j++) {
                if(i+j == hb.length) {
                    //emit log("match end",i);
                    return 0;
                }
                if(hb[i+j] != seal[j]) {
                    //emit log("not match",i);
                    break;
                }
                if(j == seal.length-1) {
                    //emit log("match ok",i);
                    return i;
                }
            }
        }
        return 0;
    }

    function cutRlpData(bytes memory pre,bytes memory hb,uint rlpSealLen,uint returnLen,uint index)public pure returns(bytes memory data){
        uint count = 0;
        uint start = 2;
        data = new bytes(returnLen);
        for (uint i=0;i<pre.length;i++){
            data[count] = pre[count];
            count++;
        }
        if (hb.length > 257){
            start = 3;
        }
        for(;start<hb.length;start++) {
            if (count>=returnLen){
                return data;
            }
            //require(count<=returnLen,"fail with cutting data.");
            if (start == index){
                data[start] = bytes1(uint8(128));
            }
            if (start>index&&start<index+rlpSealLen){
                continue;
            }else{
                data[count] = hb[start];
                count++;
            }
        }
        return data;
    }

    function cutSeal(bytes memory hb,bytes memory seal) public pure returns(bytes memory data){
        require(hb.length < 65535,"the lenght of header rlpcode is too long.");
        require(hb.length > seal.length,"params error.");

        bytes memory pre;
        uint256 len;
        (pre,len) = setPre(hb.length,seal.length);
        uint256 index = bfsearch(hb,seal);
        return cutRlpData(pre,hb,seal.length,len,index);
    }

    event log(string s,bytes b);
    function test1(bytes memory rlpHeader) public returns(bool){
        bool ret = true;
        //it only decode data about validation,so that reduce storage and calculation.
        blockHeader memory bh = decodeHeaderPart1(rlpHeader);
        //bh = decodeHeaderPart2(rlpHeader,bh);
        bytes memory extra = splitExtra(bh.extraData);
        istanbulExtra memory ist = decodeExtraData(extra);
        //the escaMsg is the hash of the header without seal by proposer and aggregated seal by validators.
        bytes memory blsMsg = cutSeal(rlpHeader,ist.seal);
        bytes memory rlpAggregatedSeal = encodeAgg(ist.aggregatedSeal.signature,ist.aggregatedSeal.round,ist.aggregatedSeal.bitmap);
        emit log("extra",extra);
        emit log("blsMsg",blsMsg);
        emit log("rlpAggregatedSeal",rlpAggregatedSeal);
        return ret;
    }

    function test2(bytes memory rlpHeader) public returns(bool){
        bool ret = true;
        //it only decode data about validation,so that reduce storage and calculation.
        blockHeader memory bh = decodeHeaderPart1(rlpHeader);
        //bh = decodeHeaderPart2(rlpHeader,bh);
        bytes memory extra = splitExtra(bh.extraData);
        istanbulExtra memory ist = decodeExtraData(extra);

        //the escaMsg is the hash of the header without seal by proposer and aggregated seal by validators.
        bytes memory blsMsg = cutSeal(rlpHeader,ist.seal);
        bytes memory rlpAggregatedSeal = encodeAgg(ist.aggregatedSeal.signature,ist.aggregatedSeal.round,ist.aggregatedSeal.bitmap);
        bytes memory ecdsaMsg = cutSeal(blsMsg,rlpAggregatedSeal);
        bytes32 ecdsaSignMsg = keccak256(abi.encodePacked(ecdsaMsg));
        //the esca seal signed by proposer
        ret = verifySign(ist.seal,ecdsaSignMsg,bh.coinbase);
        if (ret == false) {
            revert("verifyEscaSign fail");
            //return false;
        }
        return ret;
    }

}