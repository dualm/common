package common

import (
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/ugorji/go/codec"
)

type Equipment struct {
	lock        sync.Mutex
	MachineName string
	FactoryName string

	// 通信可
	commEnable bool

	lotName         string
	SubName         []string
	LoadRfid        string
	UnloadRfid      string
	PortId          string
	LotDetailJudge  string
	LeftWIPCount    float32
	LotJudge        string
	ProductType     string
	ProductionType  string
	ProductSpecName string
	ProcessFlowName string
	// step id
	ProcessOperationName string
	machineRecipeName    string
	// 工单
	ProductRequestName string
	LotGrade           string
	Node               string
	Length             string
	Location           string
	Total              uint16

	// 小母卷、子卷位置
	Position []string

	// Fields
	DvFields     map[string][]string
	SvFields     map[string][]string
	RecipeFields map[string][]string

	// alarm text
	AlarmText map[int]string

	// 缓存
	cachePath   string
	cacheHandle codec.Handle

	// 自定义字段
	customItems map[string]string
	customSet   map[string]interface{}
}

func (eqp *Equipment) Clear() error {
	eqp.lock.Lock()
	defer eqp.lock.Unlock()

	eqp.lotName = ""
	eqp.LoadRfid = ""
	eqp.ProcessOperationName = ""
	eqp.ProductRequestName = ""
	eqp.machineRecipeName = ""
	eqp.ProductType = ""
	eqp.ProductionType = ""
	eqp.ProductSpecName = ""
	eqp.SubName = make([]string, 0)

	return eqp.serialize("clear")
}

func (eqp *Equipment) SubOut(subId string) error {
	if len(eqp.SubName) == 0 {
		return fmt.Errorf("no sub products in %s", eqp.MachineName)
	}

	eqp.lock.Lock()
	defer eqp.lock.Unlock()

	var contains bool

	newSub := make([]string, 0, len(eqp.SubName)-1)

	for i := range eqp.SubName {
		if eqp.SubName[i] == subId {
			contains = true

			continue
		}

		newSub = append(newSub, eqp.SubName[i])
	}

	if !contains {
		return fmt.Errorf("sub product id not exist")
	}

	eqp.SubName = newSub

	return eqp.serialize("sub out")
}

func (eqp *Equipment) CommEnable() bool {
	return eqp.commEnable
}

func (eqp *Equipment) SetCommEnable(enable bool) error {
	eqp.lock.Lock()
	defer eqp.lock.Unlock()

	eqp.commEnable = enable

	return eqp.serialize("CommEnable")
}

func (eqp *Equipment) LotName() string {
	return eqp.lotName
}

// SetLotName, 返回的error已包含完整报错信息
func (eqp *Equipment) SetLotName(lotName string) error {
	eqp.lock.Lock()
	defer eqp.lock.Unlock()

	eqp.lotName = lotName

	return eqp.serialize("LotName")
}

func (eqp *Equipment) MachineRecipeName() string {
	return eqp.machineRecipeName
}

// SetMachineRecipeName, 返回的error已包含完整报错信息
func (eqp *Equipment) SetMachineRecipeName(recipe string) error {
	eqp.lock.Lock()
	defer eqp.lock.Unlock()

	eqp.machineRecipeName = recipe

	return eqp.serialize("MachineRecipeName")
}

func (eqp *Equipment) CustomItem(key string) string {
	if v, ok := eqp.customItems[key]; ok {
		return v
	}

	return ""
}

func (eqp *Equipment) SetCustomItem(key, value string) error {
	eqp.lock.Lock()
	defer eqp.lock.Unlock()

	if len(eqp.customItems) == 0 {
		eqp.customItems = make(map[string]string)
	}

	eqp.customItems[key] = value

	return eqp.serialize(key)
}

func (eqp *Equipment) SetCustomSet(key string, value interface{}) error {
	eqp.lock.Lock()
	defer eqp.lock.Unlock()

	if len(eqp.customSet) == 0 {
		eqp.customSet = map[string]interface{}{}
	}

	eqp.customSet[key] = value

	return eqp.serialize(key)
}

func (eqp *Equipment) CustomSet(key string) interface{} {
	if v, ok := eqp.customSet[key]; ok {
		return v
	}

	return nil
}

func (eqp *Equipment) serialize(fieldName string) error {
	_f, err := os.OpenFile(eqp.cachePath, os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}

	if err := EncodeFile(_f, eqp, eqp.cacheHandle); err != nil {
		return fmt.Errorf("CIM error when setting %s, EQP: %s, %w", fieldName, eqp.MachineName, err)
	}

	return nil
}

func NewEquipment(tmpPath, eqpName, filePattern string) (*Equipment, error) {
	if len(eqpName) == 0 {
		return nil, fmt.Errorf("make new equipment error: invalid equipment name")
	}

	// 获取缓存文件
	h := new(codec.MsgpackHandle)
	eqp, err := makeEqpFromFile(tmpPath, eqpName+filePattern, h)
	if err != nil {
		return nil, err
	}

	eqp.MachineName = eqpName

	eqp.cacheHandle = h

	return eqp, nil
}

func makeEqpFromFile(dir, pattern string, h codec.Handle) (*Equipment, error) {
	_f, err := GetFile(dir, pattern, true)
	if err != nil {
		return nil, err
	}
	defer _f.Close()

	// 解码文件
	eqp := new(Equipment)
	v, err := DecodeFile(eqp, _f, h)
	if err != nil {
		if err != io.EOF {
			return nil, fmt.Errorf("error ocurred when extracting equipment from file, error: %w", err)
		}

		eqp.cachePath = _f.Name()
		return eqp, nil
	}

	eqp = v.(*Equipment)
	eqp.cachePath = _f.Name()
	eqp.customSet = make(map[string]interface{})
	eqp.customItems = make(map[string]string)

	return eqp, nil
}
