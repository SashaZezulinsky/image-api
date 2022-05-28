package usecase

import (
	"context"
	"github.com/image-api/internal/domain"
	"github.com/image-api/internal/image/mock"
	errs "github.com/image-api/pkg/errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

var testImageHex = "ffd8ffe000104a46494600010101004800480000ffe2021c4943435f50524f46494c450001010000020c6c636d73021000006d6e74725247422058595a2007dc00010019000300290039616373704150504c0000000000000000000000000000000000000000000000000000f6d6000100000000d32d6c636d7300000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000a64657363000000fc0000005e637072740000015c0000000b777470740000016800000014626b70740000017c000000147258595a00000190000000146758595a000001a4000000146258595a000001b80000001472545243000001cc0000004067545243000001cc0000004062545243000001cc0000004064657363000000000000000363320000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000074657874000000004958000058595a20000000000000f6d6000100000000d32d58595a20000000000000031600000333000002a458595a200000000000006fa2000038f50000039058595a2000000000000062990000b785000018da58595a2000000000000024a000000f840000b6cf63757276000000000000001a000000cb01c903630592086b0bf6103f15511b3421f1299032183b92460551775ded6b707a0589b19a7cac69bf7dd3c3e930ffffffdb0084000506060709070a0b0b0a0d0e0d0e0d1312101012131d15161516151d2b1b201b1b201b2b262e2623262e264436303036444f423f424f5f55555f7872789c9cd2010506060709070a0b0b0a0d0e0d0e0d1312101012131d15161516151d2b1b201b1b201b2b262e2623262e264436303036444f423f424f5f55555f7872789c9cd2ffc2001108014c01f403012200021101031101ffc40035000002020301010100000000000000000000010206030405070809010101010101010100000000000000000000010203040506ffda000c03010002100310000000fab643edcd629ebd90709d8f242638ca30992a724e26992c64a540c824988600310c0062c73e71e2bf25db2a9e4f7bebe3ec63bbdbdbe573edc9fb7fe26fa5fb793dea719fb3e703210c131894831a9c6a2340988932a119aa836d112145259d63d7dbc49a929ab09e3753201925072e54389a009263180c06d384a452610949563f18f5df85f9f4f3adcd3ecf97dfbdd7d9d3bb55ddaaf79fd16df67f21be6b1f63e5c39be97c406d623040c4da22a492048a812424c22a459118000013663c98931e39eb59cde5f96f8cf0f47d65d9f8cf6f3bfb5f2fcd1ec3d785ea7adb3d39b72222e40c24449109800c5019184f593cbfe1bf63f0bf3fafab65e375b9fa3a30c781db538bd6e979bd3d6b2d7ecdd39fd95b5a3b9f47e0cc4da01c45b2c430837104082611538d45482230430834da58b2e24860cf1b3e2ca3fa9f97f93e8f55f572f3ef4dcfd2d093d67e86f83fadd7cbfa033f0ff64f4f93a063c972e5171201408a3789ae450c64a99b9e4b8dfcd14dbb56b87a7a393979f3dbbb589c674cf74ae76e7a3a57ee274795fa6ed9e1dee7ebf90127df845b54010c4538b0829a489248a338880a88031041834b1e4c690538a7897cc3f777c5fc3dbdab0f9cda78fababc8da9d954e7fa2e9c7977a77129b797dd172fcfff00a9fd1e4f5e754e8eb9770d1c35d0e6d7b898dd8b1d7b571ab7e9d42bf2efd175ab1c7d1c0aa5eaadbcf271dea19ede6b9ee58b59e458789bf8ef7fb7fcff00eb9e4f57b97b17987a87bbe5b907abc686426200000a13420122c104c22a484056201a8c642414e29aff002efd4fe6f8ebf2b53bd7f9de5fadc9b9d1e3d73e93d1f3cb9386ed4edd4fce3ccaf75fd397e99ed781daf38f587e57b365d955baa7575b818d37693d2ae4d56f95d28570393d3d4b3b164d3e1b567d5d5e94d72756cbd7cf5a6fa779eda8f7af5badd97d3e090ce9cd3010c13010c10d094a209a404d22340018522d71921a9063d6dc827cc3cef6df03f37d3dfaf6af567a6ad6a9696b9742bfd6a2e7967d2ed57f2eae5ae762f1d8b555baf3571c5cfc31d1cdcc899f4e3888e3e8e2d63cfad5b3999ab2e9e36e3935b6d6572aeebcbad73ea7a9966bb52ed1dfcbd168de5881881881881a65252444609312230401aae2ea42924948962a41a9e03f41f80f1f4d079b5fb1f97eae7dce1e9f5e3a1d2dec1796bd32cfc672e2f62cbca73ddeae6d75d832ea6759b0e4dfb3952d3eb4b3d7ecf0ee7a3b1bdb973cce25aa53551d5b9f396a91b4c53b7e91e7d6792fd62a4d93af3b94b166edc819098c8a9210c0100d3a435494908044317464a5639394352162a48c556b6425f8f7c93ec9f1ff009df67c47d4fceeeae96fa5dee99dbc15dddd9cbae7976fbbbac54f6bb7ccac7837b2e7a6959f63d475cfe74deb4d7f1d2755b6546ac39f86e2c2fc6e19bee3ccf33b15b69cd57b6eb9d8ae353b12583ab5feb6b37fddd1def4709010340d00c0146711000d14d0c4995118694a191324892a2488a9c48b7235bca7d7b9f9bf0aef7b072fcdece3d66c3a2ce9f53a5bb73af9795ce9ae9543b5a15b161ab5d0f55f47a6fa36f879ff008afd43e6f37e1dc5f52e4e3a79e7877ad517cb7af59a79f6be3fd2392ebd7fcc7dead6c5db91f5fcba9d9d6e7969b1d06cfa7ac6f72fa9e8f38c201a00146846982524252421a1a6688039fb387686c204c12915190c8f3fa3ce8a3503d2293e6f45579164e32475f5b919eb9a9759a773ddf6f143ad7a7c1f40db7cb3df38fb2f37cf37f42ede7d9e3f631d9e65c0f41ac677f247b879dd87f39f4fe7ce77da1cdf3f2e971b4ee7f49a75cdbe77d7f1f561abb4b86d555ed9ebb61a85bbd1e698310d00c1000081a60d038b08b71a60561c91940c21a001a10d91e7f475aa8b5cb6713cfdf8353b9f1e3cfb47d0f757c6bc93ea1d2fcd7d1f8bbdb689f487d7f8bd7ef71b37cffbbeb77aa2df3f45f1f6624ab9352bf7223e79c7eb3aedf9e742e3cdccc959cf5e881a73cf4df96182f5e5cbb0d7a25e69375f47972b4d01301000c4000009a1a60d0c88c31b4e9b8b8686a81a09b143244acd56fb47e7d385cce8f3b1bd0eb69eacb97916d9f9baf039f64e87cef4f995eed593edf836efbc5edf4c667090f5f3eb9c8afddb4d3cb38d7eae6354ec3d4e2e7a6af48a5e77b1a35fed5b6fb9532ce9eaf6aaa5a3d3e4cce2d240d50c20310040d03402600d0300c4d4b41350c64a00269a34061a45eeb535e7da9bfcfe3d75f9dd7d69791bfa1b766e74b56c773bd6cd5b46f1acb36a35b2f5633a6d6be184dedecf23b17968d4af7cdbcfcbb897be0e3a54e8decd5fc75f2edbb572a68ba716f4965ece8717d1e5be6cf9cdb379ee4b533995e39cad3043403424d20d3112880041a7403800504c4c1018474b7b1579856fd368fcbb73f02e073df475b9fdf3a973e05dba71b06de3762e5f535f3d394ab997cff63b1c49f0797bae964d4c9eafcf6c460faf9eb351f46a766e0e5dc5637e48fd4b426aa776c7dcb3271fbfafd3979b61b4d29abddb7c22dfbcfabe5e574758ca270290200010c054988861068a6260da84c6218a98045860a3df74cf20aafa5d73877f3db0f47762d570a75e3a72eaac53a829e18757b0e9e7d140b771bad8fa3de972b2f4f93be6a4ae61c7ec425862ccb9ef4b1efcd30ee2cdac98f39acf12a7e83c1dcf2bd4b75671d6dfe89e057cd67d573f2ba5be5901ca8609308920882a6018c0418c430062a6988904498414d1cea17a668e6f86cfd4e1cbb736cba1bd71d47a99359cbaf0d659e9ad39d38d2eae7cfa306d4b2de065c1a999bf2e5e4e5d36d995334f167ebc659165d41e4959af83a18b529fe7bebd41283bbab8f9f7f50bef86fa675e3747abb1acc9a201a00000000c2c2c1c640d0304a49381a608280031658c71b57a78b1ad7d99e5cdc383735e5d3c3974b9f5ccf4957465cf8d9d3d5e718dbc5338f5c9b3873eb1b3b3adb3d7965cb8f3ef9e5cf8763799c9641473474d0a9dcebb73e4f5abfd139f6e859a99d56bd9fbbe777dede7db709c31000000001886589b6b16d91241124098426008281a8d5c3b50cdc52992c31e58c69697535b1be469f5f9dc7be386ce25d4ceb3675196595c2caf2d93d8c59fa73c9b18b3ef1973e2cfacce6a54d3558399d7d367cdfcefd6bcde74ab76b87bf8eb6df53f15f41e9cbd172e9edef930250004e2022c4d35011323284355202000010c004c3142708649cb85658cb870ecc239fa3d7d7e7d39d8fa509ae664dd79ba99366766be4d8c9ace0cd9725cc333c9a8f2c7259369d116aa3833604e3f9ffa0d26bcdf53afcce7db35be93dc5f72e9d46d7dfcd980cd0224a22b01845c5ab431499000099438324270044620c709c627931e495432465c51cd8cc11cc46b476632e079d98259998de469194a54a6e41352a1a7445a8862cd8939748f40a5ea79af3fb3c3c76c1d0d29e77ea3e8be53ea5dbcfbae2ec04d01319117ffc4002410000300020202020301010100000000000102030004051112401013062030145015ffda0008010100010201f93eefe4dc886509318cbd70db1ed1f72f5e677014c5551576cfc56c3da3ee7e4dc8762734c3977d61f8e91ed1cebd9a3fe41bb3c433c66ad1db4f38c23dc269b077ff00f4d3681f48e7e43bb4c9b414b97a15588d529ed1f87ce4f937dd5ac37757f20d4e555bd063f926f05d74cf26a3e424a74d267d93f0e39b320a5c19cf678ffc8f537fbfea737b6794b430d51da8f4423239aa38cd8f64fc1cfc975e0115a4f26067afb5c5fe4494fe1df7d96bedecdb74b4c7c7d8cd09b66ac73f167f73f29d4889384693c4eb526471bce68f29df7f259af4ba6d7faedb6f67bd236612fadb5d2716d7ace919f0f3f73723b0ab58ec7fa401332bea7f9f4f738de5976c6c8a3dadb3f6792bb35efb177a568aa53ea643c63eb40ecdb8a6d31ee1ce7751e3d6a516e2d372b75d91c5ef2ec25bfd0db0f57766f27a6cd2ced9765b4315bcfab269ccc78ea70e3dde735bc5f45f5d32758e39bb39f0d6da4dc3b70dafb033d0d19ee19564f875de532853048ebbbf143493dd74e4222a43eb154d836b13f0196e8f0bab7d8bb5e4acec560126275ceea51a2c1eebc6434a80fbbcc6a5a7f64acf349d9db3eab307f10248594a2d0b09d9c1cd715c2442f03af15faf5924b3c937bac392828680ca51995e894d751aadf4f4e432d2884cc3a75387d062d210b311d6b98ba1811ee9cfca21af6b5856117d7ca47ec441adad85a99e7d5053241951a79f62b29a4ce36234d60a8d265f7b91d694370e8edc59b18263e1401a77413a2a8d8d60a30e55a140494b39e8620d71355c5691f77addd0dfd37d5d117ca22c8c67ac60eaceacebabaababb5ac27e74547b6c37383939f2cfb73115533a1c9997becbc8f1bb70d1cf32df4a2ae79b17ac1ccb8e413ded6ac183255b7efb9f8e6b58f19a3c69e3d35fe89e2bcf25ff0002d9b9c6b6a05684d5f1d89f3aac4eb2e828149eceabeb0d5e61b6341f603f0892171873ecef54c0fa8307f2ae546d6089ce9b2afb3c9a73bafb939427a583e2f1a4e89cf32eaf29c1f1da3a73e31b669e40b2a36bb6ab7a8a3f954572caea5d9f62cd5da5dee3b4b4b8fde45d62a70e6c2527ca4b57696a71d829653807523a4e3d41fceb95c64757432ae8f1dc673712df8a9e634b8cac317e0e552eb4d4871838087124d0b8185c28c5cd1c1efb8a86caafd6544b8f1cbf116e17f1de3394a4175b07c9cac845d7ceb67707b4a331a4b146881ef9cd847ca8e9846bbda32e61b7adcbe97133971f31fa1c78d9280821c7c5b624d3c96698ff008154ae57e1d42adcb1d14d7f1496ba0fd1be293ac5a56470176f6432645e19aa7bf7d86c2d314b8662518bcc41650ebe3bec904652559594cee8d2f15c0754232d8377ee9cd95b2ae74548c0925d78f58d80f7e5e5e5f15165f070e1e82aaba9029429b72b86f768365570b9c741392c1663e18f916ef637f51be09d9717faa9a74e3d75a7af04ead3a89de3b8b41ee11b3164a63d7fd52bcd22a3e5d369277fb75f514675e344698c208289255c6165aaab6b6dca80fb6c3665414c48c74f5e715f8ef0e5e19ac9df979f74596bf44678281f04512c949ab6aed4aa0fb6e97d6697dc95d774c19d67659db6315bbefcbcbe3c8ae78f5f245a568b287d5bc9c1f6e93d8d67d54d28ce6de5d9f8eddab88158119df97783e3afd3a2b78ec4ce4a9abb08dee510cba503067912492478050bd162fe40fc0c1fa75d11517cae0785356e0fb671c78850b8709efbf9ed998e0c183e460c1fab0b66ce29075ed0a8f7187407c10c0fc77e45fcfbc0060c1fa0fd88b26c65715a6da7556f6ce75f27082083f0707ea3060c18307ec728bb1af69cd818b6bd3bf68fec410430e8803aebaeba183060c1fbb0d85d8560ad27d4aafb67f720865e880bd75d75d740743f81caaecc9d5714eb522e3da3fc0823aebc7c7aebaebaebf89ca1b1aa80d913a8e3da3fc7aebaebaebaebfa1ca0be573ba644e9bafb47dd3846c9d91db64b34ca7f0fffc4003b100001030204030407070402030000000001000211032112314151226171041081911320324050a1b12330425262e1f00582c1d143a2253363ffda0008010100033f01f807a3a7e8dba9845dc7ba33aa71e41458056d1163c103c57a5ec74cec20fc5c329b9db04eed1da08cae6e84e5d029bfcca73cdcd9340b2f144b6cb86a533d47c5c51a1826e44a9b99b8d518574ec9466e95645d0bd1f6b17cedf160d692745e9fb63ef61978abcf9a9e6801dd014a21a8b3b6b5fcefefc1a983654752a93b540fba0a3d988dc226abdd37b2b652a2e6d6f153aaa43f093d562d029a8140c9115e9cea5708e9ef964294819aaaf74e2b2741e320fc9561f88a734c38f82a15858a07dc9d56bd501dc20405903d4a1126e0643728ccb81f254f629a4e6b9842e5185f6cc3fa82e11d3df2c8fa77372ba079ad829d13a775598750154a702a5f9854abb4107dc30b606a83abbfc56331cb3d86ebf2ce102c81cc3908b185c5992a4a01aa57da3026d5a1234247bee1ed4e3bdd6eac845c14ec82735bec29ca07455bb33dae64db354eac35e6fba6b8483f790809089a8b176998b5e10a6dc233767d14083f24dd3e68fe65244ba3c1334bf344ab22c713b2760783bfbee3ecd8c66d4e95c938e6e856ff0029dbca74e48e4601e8b03a43a3c156ece732e6aa3da058df23ebb4228a2b80a384dd4623fa5035b967e0163a8e911af8ec8f8c27bb92ad3c267a27f26f33754c4e6e45ee88584ec856644a145e1befa2ad07b3708d3aa438419d116bac25da21ab84a69dddf4559c738446b29e6c1be68c714273248f927d1ab8c123746b5104c48cc7735084d0a55d5820b9a2444ad015881be880a66333f40817ceca9efa211c23935536f08bb89cf3535713deee9aaecf4a6ee3e12af69f14e15321e49ed736722bed9befc5b5dc46b959705939aecd6989c7a5bb982e63ea53aa098809916baa6eca022c7e48d3765a66b51708382235523ba5b6cd459083ddaca9a87985388ea8f0b73e002dcd6191f258d8413001cd4b667053c86ee447b203749d7cd5335040cc67aa31520988f2514b16f926c12e4d73840cac154355c5c223dfb1f672f1984070927c507a2d4ed9345ac4ac7ed3a63409b111e0ac7863c5348875ff009a2735d68c1c93db026c7551aa19ca25a095c7e0ac8e20e1bdd6708b8cee8cfd512f1fcc97da61de3e6562aa5d3999f9a2f790dfcd7274098d1d0d9bccea79a7d5797ba4dbf8027619392337ccea9afa7804c6a9ad6b44648071080ab6327fda0193bfbf073483aa750ed0473405935fdce6bf2b6bcd52632093c9ad551c069fa7541a24dcac42d645b2667926ba9cb4ebe49e65b1c93b85d3d55c0574e0d6a38b0f359dd12b8fc2e809773941a5eefd2e31fceaaa398045c92210a62044cffd86bd0238604d84df9e6b03598792c203720a448db345ae9f1276fdd34d8668ce49f883831c02dfdffd252c43309d09cd58b4408c9542fb34346faa6506101bc473fdd3ea3a4dd300cefba0dcd40c405fea1706316c814c89fcc240eab8b3566ca95c5f359754034734e2622d179591da494df4752ff89a3cca740708936675d4ada30c7fd7f75f678ba01b9d5388c4c1188c2f3c93db4d98ad3a6a55e75032d02707430c3732752a954339e1b78a023e0382a3b08c8dd35fa84d69d53751f34d6b6c04f5c9716a3fb975f1129ae06dd309ff00051931c7b8c8f926f088758a77134004e5c9c364de1232176f4458e9f921aef3d161cba79a01a0cde6e8960e464225d33a29698c912d742c548e01389e11f4b6160dc23a9cca0f7c4708bf9580431971d13af19a00024e57e6aa38dbf9cbfda3910635e69f52a16926067161e2540185b0029da3bafeff005061aac3040855bd25d340e6b1d3b44eaa1f9e69b3220f24c658380e45629b4f4329ec258ee3da764d79cb169faa3fca021cc32107d29d7f928b6a715d5fc161048198b7540b663da6ab084013e08011e49c338846953701f9d0c200cca6b46632b734d8e239e4861e7a47f2cbf340faaa6cf67356db74c73ba7e18faa7f31ca53f53f2eebab7bf7a6ecee1ac590e33176eeaa173b3cd6075cd936addbe4b8507be0a87c5da4645173f47f5596190763fed3b1b49b54c9c0e4eea9ccbb321a6add9096b8019648974ce9f453fda32e6a637908fa290336dd628233cbc911131cd6161cedf44ec4458ead29f7272d112f6c8e2fa04c6e225dcbf6088f65aaa387b12554fc585a360134bb3281a936473b7705756f7fc1531b4709cd31a2730a86398e886356c937aa152f37053b1070370b132489f9a63c716994e68b4b64df476879146f7f642270d8dd09cbda328e3691a8bff847d1069cd39b4df6b8328dc1e84200a1d009b274e1e7994ca34f138c0d777218b85b655b301b3b94fc50f105178131e651cc4c744e0e10e7657081b5936d7efe11efe088289a4e8c97a2cb31aa738e764ed1e07543110f387a7d56138b3bf9a2f32e66183e69c32f10546be257f6ee34431e5fc2802235b0470ccce9e09b841fd202b29a4f1baa9f8bcc66803e37f145cfb84da62488dca7d6a924fec13fb0f66a55ab896d402083a9130982a01a2a75e8c468b0f676ce72e1e21345c6dba7372cf784d24174cf54dce3b815c3f00e129b56a4e4a9d165827bee402130ba305bcc273726fcd0e84ac36c43ea83e785b3d735a5efa14e3a212379d9024f337f28405865dd2167cd1cbf80a88b6f645b5d94f7ba70ecaf7c649e581989d8469365c361708fa21375f6ef6fff0059f920c8dd6450cac9b8a02952b87e01c3dda2e19c8abe1f9e4b08f68a71b7b3fe5359af8aa4d378246cbb3b7fe370f15d9bb43460749d41b152edec8710027d4c410ba132bff274b6000f34d77678dc2afd9de4b1a5cc26db84eab584b4c0374cecfd99d51d9009d51eea9ce7cd5e09b2e22049779c7551d5057575c3f00b77139279e688b9016a9a1864f80552b97459a36d56185fd2a8f60a15a97696567be03e9e45a62554ad2ea39b4e4aa32bb69769104e44ee8792c27d41dd49ddb38c4b5d629dd9061abc4cd2a7fb5d92bb271b7cd7f4becce92fc4e9b0172bb5f6f786fb14468852a785b164c99761b734d36b782dd348ba93aa0170fc00c774956b7d16e1025d1f242a82d71f15868dc49bdd3a9d71b294e355fb5952abd95d3673412d3b1174f7d168a9ed8027c90f5242384a352a0244c88549cd20cd976702c2dc9766a57c1279a81c9493c498ecc5829c8db92206142d9f45a205d6c95be01657455e54841ba42c74ce8ed3aac15df4dd93ce26f2dc215a99042ed8c7406621b847b3d1e2ccdca0e732803771bf419a7b7b63cc1823ce14c7ab210136c9714a846112ddfc513fa794ca7678655426e862947329ce331dc23e090567b72942404f79c74f31784ea47076963baafe926f881423d1f6665f72ab17beb567c97676f92194414f04c99839fac26477b909d50c37b842d0004d6b561e6aabaee7c72c95c12df10812adf0191dd2a02711fbaca1a8b55278e26fc976477fc43c9516461680861e484a227d79415f6467f6ff0049c778f25024fec8641cd4f9cc23b8564312b7c0a0f70cd6e119b65d1011608cebdc76214bba2dfee415cd6a9d3927066685f87cd107246660238ae14945a81f81157cd73ee139a76a447cd5a04a196bea5fd5e203bdbaa6cdd52c59b9538c953dbfea13720f70fed0a5ff00fb4f8b5079fc07e483748565872441ba0e1eff00645a50ba8c82dd020e69d7bc0e481dedba12a077de3d4a348e1325db04e7f1111cbbc260d2e9a5d1067a29d913a270d2db26ea1374503ba5614e0844211efc08c942202a8342a3aeca91b224885d07a92bb7b5e0d3730b75696dfcd39cde210477075671cef2a07784d4dc7c28c2e4ba85d0a1d3d58be48ee15c09408f7ebcf73e6caa3cf1344692ab62d4fc90162e1d0200dbd6d42bc2c23d42a466834923ee8285051408f7d908ca0df68f80506cd44e65374463d41dd64c10ed8a047ab3ddbadbee4f739a5442047bec847455bf32aee8bdb7429e9fed6fdc7d4b22e4e01391f5254f5ee9fb92102a0f748f7d943be147aa4fdcca9fb9309cb09f6548ca1429f7ebfa9afdfcfdd0d902a0f742c43e34334cd9006c148ee13f1a909bba82a1caea14b7e347ba1595d69f1ab77355d41504291f19b209ba282afddc3f1a8471774a857f8d41f56e15bee3ffc4002910010002020102050501010101000000000100112131415161102071819130a1b1c1d140e1f0f1ffda0008010100013f100f151f0a847c4843c195e43e9ac702541271cdfb4694a975bebd3da3da18663377ff008e25564af2c4db2e97154338ea85796fb910fa4fd07c2bc2bc562f120f8d43c4f2d78d4af3b111ab26870d8cd3198dabd0a3965450c5e743e846087f1ed3028777f6cad2da76888aa44462e63633cf571e2af1a95e7af2be15095e0c483c2e10f03c0f123e0794f0a9987916250c661bc4a551298d92f6c40b6af505181436e43d08d2a815abfdcc5773a1ff0065b6d74cc2e81adba8c073cc74e9eeff00955e5482315e172e0f80c1f292a5789e762ed4055f48e46b829b4e046da9c64e8ae09823697a44b40a30b5cf43ab28b52bbbb83c8aeb31d87da015037c4115d22dc08ccc3e9d4a95f55f0208a561f48e36e1c5c010017bca4a30cb1b830952bc48792e5f99421edac3751224801c55f1f10d590d506165100be8641ff652350c016d1d96094a7babf10f852a1581cca06c8683986e801d6bff00b1fc08798fa152bc6be8b7e2e264bfc3e624ea3080e7894ab0e1ca5edc4b72cac89607b90e6462f1bf5ee4168f47b3d18008d9060f854a812a54a952a54af16082c668745673709176e50e6f5965c96001ad95e8732d5d5ae1210e61e986203877feca5e125e438980cd768e5fe8097ef2d4ea20cbf2d4a95f4195e4af17c563f48e7564ab5c7785645754a953567feef155e1deabf10d7029cde7ef31a6c628a7f8c230c9d63dc95d2534e21030f13cb7e45501db386389772e1ea0ef7d6504d0ecb41957a4a000405375deb976c6b61eb799985e9cb5eec7e5bae0239facc1dc37d4100f3f89622ef7d9812be8a4af2be63c1f010e263daae5fcb8b9b0b7a828ca1d0db1d81f7599d435c39fcc40361eb490a9c1d25fb31c0ae22f14f67895498603775b86c49d4832e10f25c4901e0217010ce3703e705ab031746ef56a71831a668643ddcbed2c2e874c33d58a33643986fa37e8e256070de100dda9cb07c4b0067d215357cc50e66d57d6a586b1b10f0a952a54a95f42a54a95e43c18f8314ab132f680b251ea60ff00b08a0cb8205c8346e11b7b89a601d2bf92eb4b32559f78000e027f518c2cec1f52b99596c4b3a1d624ae8d1d1e903508b972e0066209714d8d104244b714f4a5c1c332b2b6e500de85f9ac89ebfd884b96a0e787a20a85516a0e59552e9740f58ed41aa7ee3a9ba76b65ec3fa882ca6d503e2632874330a856dd71168ba6e10f9b184aff00154a95e2c7c18e22edf9958830063a129c8ceda200b7f3ad0f4c46c192fa546d13d2a89cc3ee5f7e23aa7d51f925b5035810b8b6c0a72b1f5183f380177ac64981c545d60bb8c32d448479a98559905e2e5ccf9836de6fed323381a86ddfff004471b8e7a906c4651ffec0c50a0f44789aacdeb8edf96a165aa5bcb65b7116b6299d6525eb6683a72cd8b54d1057439fc4a862f56a077b8e0b756bfd289500d38b27d8a261555d2e5f98ae829c056e026191b58792be9d791f2b195e060bb6e56141ebc47cd5299ae4fe469ac4a295433a3e59500d1d01b7dd61ead8d5bffc236b2bdeafde618779783d486200e6c6bd96582b665bb1a8796e6c3a39c9d48784b0b96c194c341990bccc88eccc52782b130ad6db50a85db72a29ac39fec52b88cfa45a09747a9c7bc2127835d2985b010a066f27f310a9ba3d4e254f86b780ac7de15bb3952dce439fc410a636afabf677a899c39d2ac1efa22105991700ebeb031f2abfa9d6471cc39c9c0dd4c9fc43d210ff003be29122e02394e9030aa1b150eb3e485d06a51805fc7de316bba01af832c49151c203df714c143ee8f634685549ed0c10471b10ee7f1a85add4a6d4ff00c497a84acd2ea1e2a0a293d6bf72809b7d3d66202c46fac6d4be10c5501ac3ee12812c5f3f9882f8031c90b0c07ef301941486cc298ea034f31d012d5e1a05755f889900742ea3e9f98e842ad6cd64eea38d06e2a7b40e306bd480595341c6af3126d293dd5f825ba4ada6efadf58480002ba04625e585ac01ace2d6883c6f2610d4afa15fe2480158298300456eed3b47feec432dacf1333529864d033d9700650c1a00e0a8655998c30ef3357d6a22b8a0ef98188d0cd36744e4978400dedbd7b778e6e02878e9f78e46a0207bcc8f7569d16282dc8db0ca116c6085046cbd2b50ad05496dc3305149bde7711c6f4a7449602e6cad56aa50c68e535a081ee24a66af51d08582e44ed5d47e58e9d01bb39d9eef331941d8ea5d424117af207f65335556dc1d0816c0558ba0e8c6fb207758cbd6deb07ce8e68c32c5402628ec171002d2e21af25f85791fa552e109cca95e092fae37a6e895186f71ba82bd231d8f52674abb5464932905a3a5fea5d07906bbaebda32de3a1c4c8a8ee63e23b696f02607fec3a75aca4c6a7fe91401c28dda358b5df6625df158e9c3d997db6a33cdb114437637cdc7415a5cfa62e05d12d5f9c7403cc7ae2002a366bda53010f8293f93415613c1aafcdcad71459d437f046661f7eab65f4ba3bcad9d80783cf6be7d25a746ad1a72a3b422954edb4b77e874220861a7b0aaaf78ecb055761bf4c41aaa5a9c15c4afb665ece5ec5c1b9e01d3fa6122d8c2a5a43ebd7d03c08f830111d31f22c81d03ccc220f433f8866979ec7e61bda574d97a6080ce9a0b6e95cf79482b91a5fb25e1681e4cfc9986d550bc87ec5c5514851e0f559f8974432525ffef4832c4205d8caf671d20571c09defe9d92f1067865c372cb4b29dd6c43c2ceac77658c1c141db57ef0472653d72fd98118c029ea6c9911b34a66b2cbce0387be1226b144fb60f78c78ba3e97fa328d3b2d7913fe9a438421a3957d3a5ca562528383b449c29055a5c5fabc128211abce79d1fba845009a61a7e089a4cd65274fe1007763561fd8b5c8a61806a0728bc6bc2a3fe13c0f160883d60784e8ca9a8ef1a8e374e78974b190aed3e20b5b6d9c38f7fd4ca8066c673d255011e049a254dd14e988964174392e2f62475ad306b3703d1df3332bdab7cba1e89d2660bb639a4c4f62cee4165b235c8157d9fcccc597a0f5ed28128d63e10228c15677fd31467c25eda30301e08858dd5bf861a0a997b3372a10ad78dd8cdc82eafa871f310da6f6cbb7e0a89a9de749e6a5a858699579d73094dea05b5d310ae80de8b4bccb14ad815b5ae3b40196da0d1ff00658dba181815ce846d6016ec2be0992fe144281940cc57e4578d4afae781e2c48747827d539601c8e4e425e0adaff900f6f66ee262e3d999732e3ae7533533ace17bf4f599aac9af361b1ebe901af064a5e1e077da0b4ee0d3e9d5d98c00b9d42e4075e8c468a405616ed7707f90d6920df03c1d85af4a897253f9d1953003909a07ea62d115474ffdb852ac853a8d59f10d38298bc383ed2c576adb58e486a643a5729b5fb41810b51a41b49d1e5f72065f7704200306b761af5e25f05bb9d081968deee8b38d4b7eddd2e3f2445a3736264024e5bacf79921df5fe1fb97f2cb83b44abb1b8685512b152c3fe70952bc951b406cc71aeb1d5416eeb35028bcb68e2fa92bda01e3a24e213d6b88caa4a612fff005312ba285257154ca5bac51729d4ad898658830c03a76adca98c53034ed9960143ac263276f780bd12c0bc03fa4c7c46089d5f9f6955bb6c3b99cc79025a13a99128ce037eb111a286379e3d6594182c615234a184df146afb42d5a800b81aa94e544c3914b9552cedb358974534a2de0ec623738b60dc050b76dfcb1eac2f602fc592f8a6d9364ea6094b614a257bb05128c2e1304e025e8bcb164f9ad4b82eae7db43fca40f2be151d858949050f67a31d0504acdc766062876a6847b455705845aef17dc0408de82710a84e00a6ba325fde314ac0349e8e905366ae914d764e62cab6171453e9de3d52ab5c8e8ffd2140b3349abd3241f174bbd4dcb6cae5ad2d63aa60b15db10d12ad3253e60d6a21c1a5564fdc248b7407158595f31c66551ee76825142d3803f3c44b8c18ce3a1317e6721901b1a8152d3432b016d87918c716ae272f1de994e0d9ca85f7ab81026c5347ac0c4898d3f12a694c67afcb196a9def71d59b6329e90ff0be240812bcd5120127a4568a1c91405ee5b2e2add56a120e99142ce7a24d14729a7ae6e2405bc8013de1d6c53832b3e9788004b2d5b0e33cc58d0c5b2fdb9cc181a5b81fc10b4d96d2e5754900c6d0c9bc5a87a542807a2fe65442f7a6a67251e8998630dd6780e7de595862c3aafea641ef9767f12c9968eb836cb31258921ea1a95806cbbe5651ac9cb7c62335aa88ed8dfa4c70387cc5732023832fcc4f2a159569893cc8629a0f980ab659caff0025ab537a352c19b87f8abc338254a95e425786e94de3533505af12dbda3558b8dca20ef67da180a62b25be840d46978c8b683502ef955d7584544a58cc72a95dd967e1800605e03f32ef2a51dab59e92fe40217a73a3d652cf822c44c40c186110991cfa44557ea7e208ad8c1ad5a5e0154d5709a62572605f613f72b0a6787b430292b87e0f58bbaa550c677f8d42b507bc75d10e4c2934ece04220d2f1666fde63653a1785885b61bb2e2d135f625a330ff05793140f354af0666c05310c833cf104b5b1c806bd6380a981caff00f65edd28cebed9b8c8f931814f560365a9390efcc1411a7e71055900aa58fa1ac92ed27314fb32ea9a8557c0f79be96e958898757825878ef4bafd3535a42fd0e1821940e9270574d7b3002a9ce97d98ea0402b27b6bde080bae1cf1be61f20f56bde664c4b7b671880a441940d7bc2a9e077bf8095e0eb6e26b181c18829d7bdcbb09a7d7bf0af11479abc6bc0a736a855063e6188dbdd2d6ced7168fcc20cad2d1b7bb7a884b2d00c83af98d4a589e1771093210e91f4020cec003bb1e0d0061051b8315762b61bf797261d4d4838f148902d075fc952a251ea71da06e1b0993e4953a1369d5d62197de1df70eb001828d465365e14c760e2335882c2a94f488155685304308cbb5388d0087b908e9d31be65129434f0426cca57f8ebea02a0d847128c1f8dc45271c9008b6838e218858f066e3d2b42d6e86462046e5aedaf5873e91f464c26c4d31de418e4f6e26babad0e7a7b4bf7d023accbf50c3570be00d76e90d30e71ede12541059a95ca5a74541c0ef9a8e3f4f53da0d25e382700ec69674bcb2ccc545436faa930d74d65bcfbd5cac068722fdb0b32712b43a96ac68cd57f234aadeb3bf68462aa722fa10e3c95e5afa07857d4164a1b087bc56b9eb018281bc19fbc4721a3910583171c590b5e20b4b4d257318dd1b26bd655da2ea913ed139b561881ec1962527e7863603821c3019be77b97ae42bd9d38826a5c498c712cf3de562eb8d591840aa7ed16f2e72da6fa4520a2f186bd5a6a328c0f575d2a206c18c3d3e6161b4be6b11ef91f011ab42f08a7da1162a3b201f0f4601078d78d796be91f49232d40e4a8ab507017d8353442dde84d405be003efcc26915e0d63a542c126a9b6fb469555e511f0907012cc1cb4cc8f3f880a572d914b3c18a62a5053a471aa956288e2ebed1f6b2cda747ca70f42556dea4bf4c073907bdc2e03d5d7a11c699db87ed1b723669337ef303f01d20a1bb3a2648d0097097e5a95e152bfc9611df1632c39c6230b177d7985ba0bba6d22c087d4b8a2b15c2f6fdccda1be742f697299a70b4dfb759401071641b97186f1e877eb286f478810a25d350631af8588f000a6fa461cd2ba83701e1bd98fc118a818ddc517872ff00202d43bdafe21aefe1fbc341bb994590ed35257ac24b630660b0c20cbff6852e2ae96a5940a8e08d570eb82d828d63196a008240c1b1f623974c0abf64acd1ce128af7cc2855b72b6bde2095a358d400462f00f18a18859addb9f68144595b05e945c0e0b1e29fd4770ba193e77000121d4bfd90f6d0cf7feecc651d297f24ba1c7c0fdeae560a79a6dfa8688fd4dfde023111dc10230461b06e0fd065798fa8cb1144402e3b40d56d5e5b800cbd2a25caa9853bb65fe2580b13b90e952a4058723e669b7c868f88212e25caeef69b448df703a2b8edcf5e902dfe9650ef16625682d393456fde52a1d0aaf66aa2b3c3b5c2d1b0fb9f696f40ec4100cbc3c906949a4861322b0c7793de685a9509cc40a771590f35f8d7f8ae351cb9314a5ef8a978f49527ad5af965d82e5ada022abf4bfc4c8ca0cb744aa83f22c34458300e6036a364a7a81d7b451d5c70fa4a5c52ececf560906267ac15e63c91a3ab99a48de51abf5a98d6dfac1e47e3f10fafa4dc35d7ae53fa846e95d1ccac810d20d38c74615f21f241646f7fec4b16960b967854a95e77c4f2d7d04960c31d3ef1dc7c1332c63a93d6f0d17d0dce608e5a1f0400b6f526105eab0260888300ec8c5f4638bb9816bb6567aa5a5b2d323b939606dad7c7111c95dc9d1cc4665298254c501352925414a0912e87a17105232e879aa5792beba4bc0085542df489145f42fe664626f7b7fe410c0776bf10818d5f11d18a9665a3446144b16214c250fea18a7074603772b7b80e25a22bff00665564e266fe46e722bf49bd97f980e21012a54ecc46d2a662ee5ceb5280b447a170c115e52a5792bfc15e0cba2e5f794b4e00e2aae2a685ed595ec06e5268e806fdce8f4227e81a83d4b227516b9891dc13263c3c5cab0fa7881698469e8c71239e1fd3de570c747f234d4136984081025460922a2811db71d89c8324503a450065041f1bfa152be854af2d4482352c5a1efa3fb102d62f0bcb07883dd152de62cbfbc56bc04e229980b7d20d4a200801324ca54b7ae607bb87affd87cbf3e04c08103c836950e462a51547b2c0c126a5b0884b60ff008897f40468da512a2d8dedbe274fccdb19a4c402625261c457241bf069332a1f2fcc33ebf9f200812a242ad43571f88ae069f48742db28a4964a6005780f15fac7d1499a61e055bda273f1303e018acbe60e254b5ee2995e0083c0433984102102547c230a21e7941660fa22d90945809887d2a95e352a54a99879d84bf05710388c1e4bae3126906210204083ca02040878304b88e1707b406a57d63d09528c1dc74370833083cd72fc495f42a54af13c04a952a544f29fb18a628410410408208208103c8cb8845dd4154b6e33c4b83cc5da3590ec21e6af2dff008551254a89e4319c330424f0c924f0840810879197972d426dbf98c51654a17313706d300ebccb97e61f25fd310951254489e0658cb241e2c9248202040843c8f84ad9fb4758bf04b2d8fc7862a92b25e0b8a1e2bf429f35f9afc5cc211224a8913c0f99e22a04081020799a825a305ddbdaa34a9f64821cc0d121a23d489a847cc789f44f38843c18c7c2bc6a540952a10840950f2b1f0123190205ac5584784982110935108f91f0ffc4002d110001030205040103040300000000000001000211031004122021313040415105136171143281a191b1d1ffda0008010201013f0043aee7499413f60a83bbca8768f68af0aa11c2a460847ae6a1f0855726d469fb74427192504e2834b9e844edd8146cda8420e075bde0085e2ce12bf6b4fb4e73583338c2c3d5fa949aef63a035d41ba06c5a812135fed48b178f0b3945e579b86a78fba7b5b506570dbc2c3372d38f46d1d578dbf08b24ec8080a651416eb3152a6c2d014a7491b200054d995b1d39d265a4c20f9d8a0138af2a14681a7c4261dad2a54f52b9864c4c26556b9bbacc768e16f73a06a69dfb0201041f28d22065e37418e082836160363608a2f68e4c269078b8ec4b41b42dac17941116359a2bb584fafed3f23999607f8583c566c4be9fa077fc45c1dfb13c2758a63f354c83f955d82933317081caa751af6820c837217cde0b1069d3c45105f9245468e403b83f8587f9caaea593e9e679d811ff0017c7e11d4f33de21eef1e85c2f03b122c210c454c3fc83f3025b3fd1f2be4f1b4db84743812f1004af8063ff004d55ee261d5619fc0dca1628000cc6ea0691c762f17c46169560330dc7047210f87a39f339ce77db61fe9318d6b5ad020344347a1a4a3a0703a874912348b4a940cd8846c134ad8f62e0a6c2c6617d619f2c107f0aabea40c9c92a9b4b5a01327dd8a214201046426be7947ae77088dd0842e44a6b609b944a082162135de0a3d72256540a9bc58bc046a6e80dd0d06cd321475c8508a2e2107ace117fa4412804104343903bae475cdc84e088402850821a48166147ac743822106a850a3a03628f4a341d30a3a451e5378d3fffc4002c1100020201030206020202030000000000000102110310213104411220304051612271051314a13281b1ffda0008010301013f001faf8e14b4c7bb3ab86d7ef31ab95fc11e05bb216678f8b1cbe97a95a3d238ac78112c725e9638d2422097d9b2892da12bf817af8f749e8e299930a1c5ae7cf1837a222e88bb6894653fc63bb3a8c5fd79a51f87ebf4d2ec4d3127dffd964a2993c75c14f45067850a08ec2e44c72a31cfe84e58df893dfb9d6494b2f8be52f3bf4314aa688c9771c13764e290899438a28ad1b13d8b6ca462693dcc936f932e4739b7ec30494e34f9456c4d2688abfd12e0b2f4ee31f06ecbd13a252dec92dfd874c9bc895d59284a324bbd0e3777dbb12a25f4cb13d1f232425e592dbd845b524d7298ba88b9297ca1e584b95bd127b8dee329f3a4e54d086848516f81c5ae47a3f631935c0a6c72b37286b49724254599bf0c0e7f2d2163c89b978e56fef83f297470c8f66ead7ef56b6f6513b9144d2843c4fbf08c7d54252a69add24fed9922e2e98f44ce972619a9f4f9aa2a693849f0fb57ec9ff15960e9cff1f9e0ea32c1c6308ffc63fedeac7cfb14f6d149a767538965e9b14a3db9ff00c3fc779670824efc716dfc53b3f926bfb6115cc71dcbfedec3d5e4955763c52f962424568f97ec60f47c1873cf15d6e9f29f03eb9a5f8c229fc92c926e4dbde4edfdb3be95a216ad0f9f629efe47c688513c04e2d57de898a5f427f449ec34535eadf922f628689690693dc8f48de25914a325f4f75fb474d83a74e52cedf8631b496d6ccf914f239254bb2f85a444d9e2649e8a9f24a1f1ec13a6268fc877ae1c928493475599cf1469349f3fb45142b287e48b271ee85e956b5a26d33c6c6b4488a1e6fc7c3d86d0a2d8a0363f2444c9c69fb04cb1314533c07818a1bee27486c63f2c781ab4354fd92645898d9658c7e5837a4d7717b28b148723c4597e562e44c7baf4efd1b2cbf43b8b8170496fe4b3ffd9"

func TestImageUsecase_Add(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockImageRepo := mock.NewMockImageRepository(ctrl)
	authUC := NewImageUsecase(mockImageRepo)

	image := common.Hex2Bytes(testImageHex)

	returnedID := ""
	mockImageRepo.EXPECT().Add(gomock.Any(), "", gomock.Any(), domain.Image(image), map[string]interface{}{"height": "332", "image_type": "image/jpeg", "width": "500"}).Return(returnedID, nil)

	id, err := authUC.Add(context.Background(), domain.Image(image))

	require.NoError(t, err)
	require.NotNil(t, id)
	require.Equal(t, returnedID, id)
}

func TestImageUsecase_AddWrongImageData(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockImageRepo := mock.NewMockImageRepository(ctrl)
	authUC := NewImageUsecase(mockImageRepo)

	image := common.Hex2Bytes("123")

	_, err := authUC.Add(context.Background(), domain.Image(image))

	require.NotNil(t, err)
	require.Equal(t, err, errs.ErrFormat)
}

func TestImageUsecase_Get(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockImageRepo := mock.NewMockImageRepository(ctrl)
	authUC := NewImageUsecase(mockImageRepo)

	image := common.Hex2Bytes(testImageHex)

	objID := "6290f67ea73fa0eddf89a738"
	mockImageRepo.EXPECT().Get(gomock.Any(), objID).Return(image, nil)

	returnedImage, err := authUC.Get(context.Background(), objID)

	require.NoError(t, err)
	require.NotNil(t, returnedImage)
	require.Equal(t, domain.Image(image), returnedImage)
}

func TestImageUsecase_GetWrongID(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockImageRepo := mock.NewMockImageRepository(ctrl)
	authUC := NewImageUsecase(mockImageRepo)

	objID := "bad"
	mockImageRepo.EXPECT().Get(gomock.Any(), objID).Return(nil, errs.ErrNotFound)

	returnedImage, err := authUC.Get(context.Background(), objID)

	require.Equal(t, errs.ErrNotFound, err)
	require.Equal(t, domain.Image(nil), returnedImage)
}

func TestImageUsecase_GetMetadata(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockImageRepo := mock.NewMockImageRepository(ctrl)
	authUC := NewImageUsecase(mockImageRepo)

	metadata := map[string]interface{}{
		"height":     "2560",
		"image_type": "image/png",
		"width":      "4096",
	}
	objID := "6290f67ea73fa0eddf89a738"

	mockImageRepo.EXPECT().GetMetadata(gomock.Any(), objID).Return(metadata, nil)

	returnedMetadata, err := authUC.GetMetadata(context.Background(), objID)

	require.NoError(t, err)
	require.NotNil(t, returnedMetadata)
	require.Equal(t, metadata, returnedMetadata)
}

func TestImageUsecase_GetMetadataWrongID(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockImageRepo := mock.NewMockImageRepository(ctrl)
	authUC := NewImageUsecase(mockImageRepo)

	objID := "bad"

	mockImageRepo.EXPECT().GetMetadata(gomock.Any(), objID).Return(nil, errs.ErrNotFound)

	returnedMetadata, err := authUC.GetMetadata(context.Background(), objID)

	require.Nil(t, returnedMetadata)
	require.NotNil(t, err)
	require.Equal(t, errs.ErrNotFound, err)
}

func TestImageUsecase_ListAllMetadata(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockImageRepo := mock.NewMockImageRepository(ctrl)
	authUC := NewImageUsecase(mockImageRepo)

	metadata := []map[string]interface{}{
		{
			"ChunkSize":  261120,
			"ID":         "629102afd39331648b1f8b57",
			"Length":     14368,
			"Metadata":   "{\"height\": \"332\",\"width\": \"500\",\"image_type\": \"image/jpeg\"}",
			"Name":       "de3d2e949a8acca0f51aa2ba8f2033a7f685af4fea5df9e2",
			"UploadDate": "2022-05-27T16:56:15.809Z",
		},
		{
			"ChunkSize":  261120,
			"ID":         "629102ce733bc6107f8d5c2e",
			"Length":     14368,
			"Metadata":   "{\"image_type\": \"image/jpeg\",\"height\": \"332\",\"width\": \"500\"}",
			"Name":       "d075769cabe189d5530c38bf727e2d4367bb39b9b6d56297",
			"UploadDate": "2022-05-27T16:56:46.334Z",
		},
	}
	mockImageRepo.EXPECT().ListAllMetadata(gomock.Any()).Return(metadata, nil)

	returnedMetadata, err := authUC.ListAllMetadata(context.Background())

	require.NoError(t, err)
	require.NotNil(t, returnedMetadata)
	require.Equal(t, metadata, returnedMetadata)
}

func TestImageUsecase_ListAllMetadataNoMetadata(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockImageRepo := mock.NewMockImageRepository(ctrl)
	authUC := NewImageUsecase(mockImageRepo)

	metadata := []map[string]interface{}{}
	mockImageRepo.EXPECT().ListAllMetadata(gomock.Any()).Return(metadata, nil)

	returnedMetadata, err := authUC.ListAllMetadata(context.Background())

	require.NoError(t, err)
	require.NotNil(t, returnedMetadata)
	require.Equal(t, metadata, returnedMetadata)
}

func TestImageUsecase_Update(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockImageRepo := mock.NewMockImageRepository(ctrl)
	authUC := NewImageUsecase(mockImageRepo)

	objID := "6290f67ea73fa0eddf89a738"
	image := common.Hex2Bytes(testImageHex)
	metadata := map[string]interface{}{
		"height":     "332",
		"image_type": "image/jpeg",
		"width":      "500",
	}

	mockImageRepo.EXPECT().Delete(gomock.Any(), objID).Return(nil)
	mockImageRepo.EXPECT().Add(gomock.Any(), objID, gomock.Any(), domain.Image(image), metadata).Return(objID, nil)

	err := authUC.Update(context.Background(), objID, image)

	require.NoError(t, err)
}

func TestImageUsecase_UpdateWrongID(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockImageRepo := mock.NewMockImageRepository(ctrl)
	authUC := NewImageUsecase(mockImageRepo)

	objID := "bad"
	image := common.Hex2Bytes(testImageHex)

	mockImageRepo.EXPECT().Delete(gomock.Any(), objID).Return(errs.ErrBadID)

	err := authUC.Update(context.Background(), objID, image)

	require.Equal(t, errs.ErrBadID, err)
}

func TestImageUsecase_UpdateWrongImage(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockImageRepo := mock.NewMockImageRepository(ctrl)
	authUC := NewImageUsecase(mockImageRepo)

	objID := "6290f67ea73fa0eddf89a738"

	mockImageRepo.EXPECT().Delete(gomock.Any(), objID).Return(nil)

	err := authUC.Update(context.Background(), objID, domain.Image("bad"))

	require.NotNil(t, err)
	require.Equal(t, errs.ErrFormat, err)
}
