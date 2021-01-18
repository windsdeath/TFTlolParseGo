package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	//"reflect"
)

type ChampElement struct {
	Kind []string `json:"kind"`
	Job  []string `json:"job"`
	Name string   `json:"name"`
	Cost string   `json:"cost"`
}

var ChosenChamps []string

func find(ChampNames []string, Level int, Chosen string, champs []ChampElement) []ChampElement {
	// ChampNames에 있는 챔프의 정보를 ChosenChamps에 넣는다 (name, kind, job, cost)
	var finded []ChampElement
	for _, champ := range champs { // champs중 ChampNames의 값과 일치하는 값의 구조체를 반환
		for _, champName := range ChampNames {
			if champName == champ.Name {
				finded = append(finded, champ)
			}
		}
	}
	return finded
}

func main() {
	doc := `
[{
    "kind": ["대장군"],
    "job": ["선봉대"],
    "name": "가렌",
    "cost": "1"
}, {
    "kind": ["신성"],
    "job": ["흡수자"],
    "name": "나서스",
    "cost": "1"
}, {
    "kind": ["우화"],
    "job": ["선봉대"],
    "name": "노틸러스",
    "cost": "2"
}, {
    "kind": ["나무정령"],
    "job": ["싸움꾼"],
    "name": "누누와 윌럼프",
    "cost": "3"
}, {
    "kind": ["대장군"],
    "job": ["명사수"],
    "name": "니달리",
    "cost": "1"
}, {
    "kind": ["우화"],
    "job": ["신비술사"],
    "name": "니코",
    "cost": "3"
}, {
    "kind": ["행운"],
    "job": ["학살자"],
    "name": "다리우스",
    "cost": "3"
}, {
    "kind": ["영혼"],
    "job": ["암살자"],
    "name": "다이애나",
    "cost": "1"
}, {
    "kind": ["나무정령"],
    "job": ["귀감"],
    "name": "라칸",
    "cost": "2"
}, {
    "kind": ["나무정령"],
    "job": ["요술사"],
    "name": "룰루",
    "cost": "2"
}, {
    "kind": ["신성"],
    "job": ["결투가"],
    "name": "리신",
    "cost": "5"
}, {
    "kind": ["나무정령"],
    "job": ["싸움꾼"],
    "name": "마오카이",
    "cost": "1"
}, {
    "kind": ["선지자"],
    "job": ["흡수자"],
    "name": "모르가나",
    "cost": "4"
}, {
    "kind": ["대장군"],
    "job": ["싸움꾼"],
    "name": "바이",
    "cost": "2"
}, {
    "kind": ["나무정령"],
    "job": ["요술사"],
    "name": "베이가",
    "cost": "3"
}, {
    "kind": ["용의 영혼"],
    "job": ["선봉대"],
    "name": "브라움",
    "cost": "2"
}, {
    "kind": ["용의 영혼"],
    "job": ["요술사"],
    "name": "브랜드",
    "cost": "1"
}, {
    "kind": ["광신도"],
    "job": ["흡수자"],
    "name": "블라디미르",
    "cost": "2"
}, {
    "kind": ["무모한 자"],
    "job": ["명사수", "학살자"],
    "name": "사미라",
    "cost": "5"
}, {
    "kind": ["행운"],
    "job": ["선봉대"],
    "name": "세주아니",
    "cost": "4"
}, {
    "kind": ["우두머리"],
    "job": ["싸움꾼"],
    "name": "세트",
    "cost": "5"
}, {
    "kind": ["닌자"],
    "job": ["신비술사"],
    "name": "쉔",
    "cost": "4"
}, {
    "kind": ["용의 영혼"],
    "job": ["싸움꾼"],
    "name": "쉬바나",
    "cost": "3"
}, {
    "kind": ["용의 영혼"],
    "job": ["흡수자"],
    "name": "스웨인",
    "cost": "5"
}, {
    "kind": ["광신도"],
    "job": ["명사수"],
    "name": "시비르",
    "cost": "3"
}, {
    "kind": ["용의 영혼"],
    "job": ["요술사"],
    "name": "아우렐리온 솔",
    "cost": "4"
}, {
    "kind": ["대장군"],
    "job": ["귀감", "황제"],
    "name": "아지르",
    "cost": "5"
}, {
    "kind": ["닌자"],
    "job": ["암살자"],
    "name": "아칼리",
    "cost": "3"
}, {
    "kind": ["광신도"],
    "job": ["선봉대"],
    "name": "아트록스",
    "cost": "4"
}, {
    "kind": ["행운"],
    "job": ["요술사"],
    "name": "애니",
    "cost": "2"
}, {
    "kind": ["추방자"],
    "job": ["결투가"],
    "name": "야스오",
    "cost": "1"
}, {
    "kind": ["광신도"],
    "job": ["귀감"],
    "name": "엘리스",
    "cost": "1"
}, {
    "kind": ["신성"],
    "job": ["선봉대"],
    "name": "오공",
    "cost": "1"
}, {
    "kind": ["나무정령"],
    "job": ["대장장이", "선봉대"],
    "name": "오른",
    "cost": "5"
}, {
    "kind": ["용의 영혼"],
    "job": ["학살자"],
    "name": "올라프",
    "cost": "4"
}, {
    "kind": ["추방자"],
    "job": ["조율자"],
    "name": "요네",
    "cost": "5"
}, {
    "kind": ["영혼"],
    "job": ["신비술사"],
    "name": "유미",
    "cost": "3"
}, {
    "kind": ["선지자", "신성"],
    "job": ["조율자"],
    "name": "이렐리아",
    "cost": "3"
}, {
    "kind": ["대장군"],
    "job": ["귀감"],
    "name": "자르반 4세",
    "cost": "4"
}, {
    "kind": ["나무정령"],
    "job": ["귀감", "처형자"],
    "name": "자야",
    "cost": "4"
}, {
    "kind": ["선지자"],
    "job": ["신비술사"],
    "name": "잔나",
    "cost": "2"
}, {
    "kind": ["신성"],
    "job": ["결투가"],
    "name": "잭스",
    "cost": "2"
}, {
    "kind": ["닌자"],
    "job": ["학살자"],
    "name": "제드",
    "cost": "2"
}, {
    "kind": ["광신도"],
    "job": ["신비술사"],
    "name": "질리언",
    "cost": "5"
}, {
    "kind": ["우화"],
    "job": ["싸움꾼"],
    "name": "초가스",
    "cost": "4"
}, {
    "kind": ["대장군", "행운"],
    "job": ["암살자"],
    "name": "카타리나",
    "cost": "3"
}, {
    "kind": ["광신도"],
    "job": ["결투가"],
    "name": "칼리스타",
    "cost": "3"
}, {
    "kind": ["닌자"],
    "job": ["귀감"],
    "name": "케넨",
    "cost": "3"
}, {
    "kind": ["신성"],
    "job": ["처형자"],
    "name": "케일",
    "cost": "4"
}, {
    "kind": ["영혼"],
    "job": ["처형자"],
    "name": "킨드레드",
    "cost": "3"
}, {
    "kind": ["선지자"],
    "job": ["암살자"],
    "name": "탈론",
    "cost": "4"
}, {
    "kind": ["행운"],
    "job": ["싸움꾼"],
    "name": "탐 켄치",
    "cost": "1"
}, {
    "kind": ["용의 영혼"],
    "job": ["명사수"],
    "name": "트리스타나",
    "cost": "1"
}, {
    "kind": ["대장군"],
    "job": ["결투가", "학살자"],
    "name": "트린다미어",
    "cost": "4"
}, {
    "kind": ["광신도"],
    "job": ["요술사"],
    "name": "트위스티드 페이트",
    "cost": "1"
}, {
    "kind": ["영혼"],
    "job": ["명사수"],
    "name": "티모",
    "cost": "2"
}, {
    "kind": ["광신도"],
    "job": ["암살자", "학살자"],
    "name": "파이크",
    "cost": "2"
}, {
    "kind": ["선지자"],
    "job": ["결투가"],
    "name": "피오라",
    "cost": "1"
}]
`
	var champs []ChampElement
	json.Unmarshal([]byte(doc), &champs)

	champ, err := ioutil.ReadFile("champ.json")
	if err != nil {
		return
	}
	json.Unmarshal(champ, &champs)
	//fmt.Print(champs[0].Job[0])
	finded := find([]string{"피오라", "파이크"}, 5, "hi", champs)
	fmt.Println(finded)
}

// ChampNames에 있는 챔프를 ChampList에서 삭제
// ChosenChamps의 test(ChosenChamp.job)+ test(ChosenChamp.kind)+ cost +Chosen했을때의 값이 높은 것 우선으로 +2했을때 +3했을때 총 밸류가 높은 값을 choice에 저장
// 위 함수를 목표Level까지 재귀실행
// 결과값 choice를 반환

// func countValue(choosenChamp []ChampElement, choosen string) { // 골라진 챔프와 선받자를 받아 현재 밸류를 리턴
// 	let sumCostValue = choosenChamp.reduce(function(acc, cc) {
// 		return Number(acc.cost) + Number(cc.cost)
// 	})
// 	let objChamps = makeSynergyObj(choosenChamp)
// 	let sumSynergy = function(objChamps, choosen) {
// 		objChamps[choosen] += 1;
// 		currentValue = 0;

// 		switch objChamps {
// 			case objChamps["광신도"] >= 3:
// 				currentValue += 4
// 			case objChamps["광신도"] >= 6:
// 				currentValue += 7
// 			case objChamps["광신도"] >= 9:
// 				currentValue += 11
// 			case objChamps["나무정령"] >= 3:
// 				currentValue += 4
// 			case objChamps["나무정령"] >= 6:
// 				currentValue += 9
// 			case objChamps["나무정령"] >= 9:
// 				currentValue += 12
// 			case objChamps["닌자"] = 1:
// 				currentValue += 2
// 			case objChamps["닌자"] = 4:
// 				currentValue += 10
// 			case objChamps["대장군"] >= 3:
// 				currentValue += 4
// 			case objChamps["대장군"] >= 6:
// 				currentValue += 10
// 			case objChamps["대장군"] >= 9:
// 				currentValue += 10
// 			case objChamps["무모한 자"] >= 1:
// 				currentValue += 4
// 			case objChamps["선지자"] >= 2:
// 				currentValue += 2
// 			case objChamps["선지자"] >= 4:
// 				currentValue += 6
// 			case objChamps["선지자"] >= 6:
// 				currentValue += 9
// 			case objChamps["신성"] >= 2:
// 				currentValue += 3
// 			case objChamps["신성"] >= 4:
// 				currentValue += 4
// 			case objChamps["신성"] >= 6:
// 				currentValue += 7
// 			case objChamps["신성"] >= 8:
// 				currentValue += 10
// 			case objChamps["영혼"] >= 2:
// 				currentValue += 4
// 			case objChamps["영혼"] >= 4:
// 				currentValue += 10
// 			case objChamps["용의 영혼"] >= 3:
// 				currentValue += 4
// 			case objChamps["용의 영혼"] >= 6:
// 				currentValue += 11
// 			case objChamps["용의 영혼"] >= 9:
// 				currentValue += 8
// 			case objChamps["우두머리"] >= 1:
// 				currentValue += 4
// 			case objChamps["우화"] >= 3:
// 				currentValue += 10
// 			case objChamps["추방자"] >= 1:
// 				currentValue += 1
// 			case objChamps["추방자"] >= 2:
// 				currentValue += 6
// 			case objChamps["행운"] >= 3:
// 				currentValue += 6
// 			case objChamps["행운"] >= 6:
// 				currentValue += 12
// 			case objChamps["결투가"] >= 2:
// 				currentValue += 3
// 			case objChamps["결투가"] >= 4:
// 				currentValue += 4
// 			case objChamps["결투가"] >= 6:
// 				currentValue += 8
// 			case objChamps["결투가"] >= 8:
// 				currentValue += 9
// 			case objChamps["귀감"] >= 2:
// 				currentValue += 4
// 			case objChamps["귀감"] >= 4:
// 				currentValue += 6
// 			case objChamps["귀감"] >= 6:
// 				currentValue += 9
// 			case objChamps["대장장이"] = 1:
// 				currentValue += 10
// 			case objChamps["명사수"] >= 2:
// 				currentValue += 4
// 			case objChamps["명사수"] >= 4:
// 				currentValue += 7
// 			case objChamps["명사수"] >= 6:
// 				currentValue += 10
// 			case objChamps["선봉대"] >= 2:
// 				currentValue += 3
// 			case objChamps["선봉대"] >= 4:
// 				currentValue += 5
// 			case objChamps["선봉대"] >= 6:
// 				currentValue += 7
// 			case objChamps["선봉대"] >= 8:
// 				currentValue += 9
// 			case objChamps["신비술사"] >= 2:
// 				currentValue += 4
// 			case objChamps["신비술사"] >= 4:
// 				currentValue += 7
// 			case objChamps["신비술사"] >= 6:
// 				currentValue += 9
// 			case objChamps["싸움꾼"] >= 2:
// 				currentValue += 3
// 			case objChamps["싸움꾼"] >= 4:
// 				currentValue += 5
// 			case objChamps["싸움꾼"] >= 6:
// 				currentValue += 6
// 			case objChamps["싸움꾼"] >= 8:
// 				currentValue += 9
// 			case objChamps["암살자"] >= 2:
// 				currentValue += 3
// 			case objChamps["암살자"] >= 4:
// 				currentValue += 6
// 			case objChamps["암살자"] >= 6:
// 				currentValue += 10
// 			case objChamps["요술사"] >= 3:
// 				currentValue += 4
// 			case objChamps["요술사"] >= 5:
// 				currentValue += 7
// 			case objChamps["요술사"] >= 7:
// 				currentValue += 11
// 			case objChamps["조율자"] >= 2:
// 				currentValue += 2
// 			case objChamps["조율자"] >= 3:
// 				currentValue += 4
// 			case objChamps["조율자"] >= 4:
// 				currentValue += 5
// 			case objChamps["처형자"] >= 2:
// 				currentValue += 2
// 			case objChamps["처형자"] >= 3:
// 				currentValue += 4
// 			case objChamps["처형자"] >= 4:
// 				currentValue += 5
// 			case objChamps["학살자"] >= 3:
// 				currentValue += 5
// 			case objChamps["학살자"] >= 6:
// 				currentValue += 12
// 			case objChamps["황제"] >= 1:
// 				currentValue += 5
// 			case objChamps["흡수자"] >= 2:
// 				currentValue += 4
// 			case objChamps["흡수자"] >= 4:
// 				currentValue += 10
// 		}
// 		return currentValue;
// 	}
// 	let sumSynergyValue = sumSynergy(objChamps, choosen)
// 	let currentValue = sumCostValue + sumSynergyValue;
// 	return currentValue;
// }

// func main() {

// }

// function joinChamps(choosenChamp) {
//     return choosenChamp.reduce(function(arr, val) { // ["요술사","귀감","황제"]
//         return arr.concat(val.job.map((x, i) => x));
//     }, [])
// }

// function makeSynergyObj(choosenChamps) { // 챔프들의 배열을 받아 시너지값의 객체로 리턴함
//     return choosenChamps.reduce(function(SynergyObj, synergy) { // {요술사:1, 귀감:1, 황제:1}
//         if (synergy in SynergyObj) {
//             SynergyObj[synergy]++;
//         } else {
//             SynergyObj[synergy] = 1;
//         }
//         return SynergyObj;
//     }, {})
// }

// function choosenChamps() { // 챔프이름의 배열을 받아 챔프정보가 들어간 배열로 리턴
//     let copyChamps = champs
//     let choosenChamps = []
//     for (i in arguments[0]) {
//         let index = champs.findIndex(x => x.name == arguments[0][i]);
//         choosenChamps.push(copyChamps.splice(index, 1)[0])
//     }
//     return choosenChamps
// }

// function popChamp(openChamps, choosenC) {
//     for (c in choosenC) {
//         openChamps.pop(c);
//     }
// }

// const dfs = (openChamp, level, arr = []) => {
//     //3개를 선택한다는가정에 3개가 선택 됐다면 출력
//     if (level === level - openChamp.length) info.push([...arr]);
//     else {
//         for (let i = 0; i < openChamp.length; i++) {
//             arr.push(openChamp[i]);
//             dfs(openChamp.slice(i + 1), level + 1, arr);
//             arr.pop();
//         }
//     }
// };

// let info = [];

// function choice(choosenChamp, openChamps, level, choosen) { // 더 탐색할 챔프들의 배열을 받아
//     let maxValue = 0;
//     if (level === 1) {
//         return {
//             "maxValue": choosenChamp["cost"] + sumSynergy(makeSynergyObj(choosenChamp)),
//             "Champs": choosenChamp["name"],
//             "choosen": choosen
//         }
//     } //레벨이 1이면 선택된 챔프의 정보로 바로 반환
//     let results = dfs(openChamps, level);

//     results.forEach(
//         function(x) {
//             let currentValue = countValue(x);
//             if (maxValue < currentValue) {
//                 maxValue = currentValue;
//                 info = [{ "maxValue": maxValue, "Champs": choosenChamp, "choosen": choosen }];
//             } else if (maxValue == currentValue) {
//                 info.push({ "maxValue": maxValue, "Champs": choosenChamp, "choosen": choosen });
//             }
//         }
//     )
//     return info
// }

// function find(ChampNamesArray, level, choosen) { //챔프와 레벨을 받아서 최적의 조합 리턴
//     if (!Array.isArray(ChampNamesArray)) {
//         return console.log("챔프이름은 배열로만 입력해주세요.")
//     } else if (typeof level !== "number") {
//         return console.log("레벨은 숫자로만 입력해주세요.")
//     } else if (typeof choosen !== "string") {
//         return console.log("선받자는 문자열로만 입력해주세요.")
//     }
//     let choosenC = choosenChamps(ChampNamesArray); // ChampNamesArray을 받아 choosenC 배열에 챔프정보 저장
//     let openChamps = [...champs]; //챔프들 복사, 더 탐색해볼 챔프의 배열
//     popChamp(openChamps, choosenC); //openChams 배열에서 선택된 챔프를 뺌
//     // choosenC 는 arg로 받은 챔프들을 가지고있고, openChamps는 arg로 받은 챔프 외의 챔프만 가지고 있는 상태
//     console.log(choice(choosenC, openChamps, level, choosen));
// }
// find(["다리우스", "노틸러스"], 3, "우화");
