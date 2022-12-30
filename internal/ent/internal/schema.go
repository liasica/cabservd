// Code generated by ent, DO NOT EDIT.

//go:build tools
// +build tools

// Package internal holds a loadable version of the latest schema.
package internal

const Schema = `{"Schema":"github.com/auroraride/cabservd/internal/ent/schema","Package":"github.com/auroraride/cabservd/internal/ent","Schemas":[{"name":"Bin","config":{"Table":""},"edges":[{"name":"cabinet","type":"Cabinet","field":"cabinet_id","ref_name":"bins","unique":true,"inverse":true,"required":true}],"fields":[{"name":"created_at","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"default":true,"default_kind":19,"immutable":true,"position":{"Index":0,"MixedIn":true,"MixinIndex":0}},{"name":"updated_at","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"default":true,"default_kind":19,"update_default":true,"position":{"Index":1,"MixedIn":true,"MixinIndex":0}},{"name":"uuid","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"size":32,"unique":true,"validators":1,"position":{"Index":0,"MixedIn":false,"MixinIndex":0},"comment":"唯一标识"},{"name":"cabinet_id","type":{"Type":18,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":1,"MixedIn":false,"MixinIndex":0}},{"name":"brand","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":2,"MixedIn":false,"MixinIndex":0},"comment":"品牌"},{"name":"serial","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":3,"MixedIn":false,"MixinIndex":0},"comment":"电柜设备序列号"},{"name":"name","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":4,"MixedIn":false,"MixinIndex":0},"comment":"仓位名称(N号仓)"},{"name":"ordinal","type":{"Type":12,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":5,"MixedIn":false,"MixinIndex":0},"comment":"仓位序号(从1开始)"},{"name":"open","type":{"Type":1,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"default":true,"default_value":false,"default_kind":1,"position":{"Index":6,"MixedIn":false,"MixinIndex":0},"comment":"仓门是否开启"},{"name":"enable","type":{"Type":1,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"default":true,"default_value":true,"default_kind":1,"position":{"Index":7,"MixedIn":false,"MixinIndex":0},"comment":"仓位是否启用"},{"name":"health","type":{"Type":1,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"default":true,"default_value":true,"default_kind":1,"position":{"Index":8,"MixedIn":false,"MixinIndex":0},"comment":"仓位是否健康"},{"name":"battery_exists","type":{"Type":1,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"default":true,"default_value":false,"default_kind":1,"position":{"Index":9,"MixedIn":false,"MixinIndex":0},"comment":"是否有电池"},{"name":"battery_sn","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"default":true,"default_value":"","default_kind":24,"position":{"Index":10,"MixedIn":false,"MixinIndex":0},"comment":"电池序列号"},{"name":"voltage","type":{"Type":20,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"default":true,"default_value":0,"default_kind":14,"position":{"Index":11,"MixedIn":false,"MixinIndex":0},"comment":"当前电压"},{"name":"current","type":{"Type":20,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"default":true,"default_value":0,"default_kind":14,"position":{"Index":12,"MixedIn":false,"MixinIndex":0},"comment":"当前电流"},{"name":"soc","type":{"Type":20,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"default":true,"default_value":0,"default_kind":14,"position":{"Index":13,"MixedIn":false,"MixinIndex":0},"comment":"电池电量"},{"name":"soh","type":{"Type":20,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"default":true,"default_value":0,"default_kind":14,"position":{"Index":14,"MixedIn":false,"MixinIndex":0},"comment":"电池健康程度"},{"name":"remark","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"nillable":true,"optional":true,"position":{"Index":15,"MixedIn":false,"MixinIndex":0},"comment":"仓位备注"}],"indexes":[{"fields":["created_at"]},{"fields":["cabinet_id"]},{"fields":["serial","brand"]},{"fields":["battery_exists"]},{"fields":["ordinal"]},{"fields":["battery_sn"]},{"fields":["soc"]}],"annotations":{"EntSQL":{"table":"bin","with_comments":true}}},{"name":"Cabinet","config":{"Table":""},"edges":[{"name":"bins","type":"Bin"}],"fields":[{"name":"created_at","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"default":true,"default_kind":19,"immutable":true,"position":{"Index":0,"MixedIn":true,"MixinIndex":0}},{"name":"updated_at","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"default":true,"default_kind":19,"update_default":true,"position":{"Index":1,"MixedIn":true,"MixinIndex":0}},{"name":"online","type":{"Type":1,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"default":true,"default_value":false,"default_kind":1,"position":{"Index":0,"MixedIn":false,"MixinIndex":0},"comment":"是否在线"},{"name":"brand","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":1,"MixedIn":false,"MixinIndex":0},"comment":"品牌"},{"name":"serial","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"unique":true,"position":{"Index":2,"MixedIn":false,"MixinIndex":0},"comment":"电柜编号"},{"name":"status","type":{"Type":6,"Ident":"cabinet.Status","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"enums":[{"N":"initializing","V":"initializing"},{"N":"idle","V":"idle"},{"N":"busy","V":"busy"},{"N":"exchange","V":"exchange"},{"N":"abnormal","V":"abnormal"}],"default":true,"default_value":"initializing","default_kind":24,"position":{"Index":3,"MixedIn":false,"MixinIndex":0},"comment":"状态"},{"name":"enable","type":{"Type":1,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"default":true,"default_value":false,"default_kind":1,"position":{"Index":4,"MixedIn":false,"MixinIndex":0},"comment":"电柜是否启用"},{"name":"lng","type":{"Type":20,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"nillable":true,"optional":true,"position":{"Index":5,"MixedIn":false,"MixinIndex":0},"comment":"经度"},{"name":"lat","type":{"Type":20,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"nillable":true,"optional":true,"position":{"Index":6,"MixedIn":false,"MixinIndex":0},"comment":"纬度"},{"name":"gsm","type":{"Type":20,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"nillable":true,"optional":true,"position":{"Index":7,"MixedIn":false,"MixinIndex":0},"comment":"GSM信号强度"},{"name":"voltage","type":{"Type":20,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"nillable":true,"optional":true,"position":{"Index":8,"MixedIn":false,"MixinIndex":0},"comment":"换电柜总电压 (V)"},{"name":"current","type":{"Type":20,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"nillable":true,"optional":true,"position":{"Index":9,"MixedIn":false,"MixinIndex":0},"comment":"换电柜总电流 (A)"},{"name":"temperature","type":{"Type":20,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"nillable":true,"optional":true,"position":{"Index":10,"MixedIn":false,"MixinIndex":0},"comment":"柜体温度值 (换电柜温度)"},{"name":"electricity","type":{"Type":20,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"nillable":true,"optional":true,"position":{"Index":11,"MixedIn":false,"MixinIndex":0},"comment":"总用电量"}],"indexes":[{"fields":["created_at"]},{"fields":["brand"]},{"fields":["status"]},{"fields":["enable"]},{"fields":["lng"]},{"fields":["lat"]}],"annotations":{"EntSQL":{"table":"cabinet","with_comments":true}}},{"name":"Console","config":{"Table":""},"edges":[{"name":"cabinet","type":"Cabinet","field":"cabinet_id","unique":true,"required":true},{"name":"bin","type":"Bin","field":"bin_id","unique":true,"required":true}],"fields":[{"name":"cabinet_id","type":{"Type":18,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":0,"MixedIn":true,"MixinIndex":0}},{"name":"bin_id","type":{"Type":18,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":0,"MixedIn":true,"MixinIndex":1}},{"name":"operate","type":{"Type":8,"Ident":"model.OperateType","PkgPath":"github.com/auroraride/adapter/model","PkgName":"model","Nillable":false,"RType":{"Name":"OperateType","Ident":"model.OperateType","Kind":24,"PkgPath":"github.com/auroraride/adapter/model","Methods":{"Scan":{"In":[{"Name":"","Ident":"interface {}","Kind":20,"PkgPath":"","Methods":null}],"Out":[{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"Value":{"In":[],"Out":[{"Name":"Value","Ident":"driver.Value","Kind":20,"PkgPath":"database/sql/driver","Methods":null},{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]}}}},"nillable":true,"optional":true,"position":{"Index":0,"MixedIn":false,"MixinIndex":0},"schema_type":{"postgres":"varchar"},"comment":"操作"},{"name":"serial","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":1,"MixedIn":false,"MixinIndex":0},"comment":"电柜设备序列号"},{"name":"uuid","type":{"Type":4,"Ident":"uuid.UUID","PkgPath":"github.com/google/uuid","PkgName":"uuid","Nillable":false,"RType":{"Name":"UUID","Ident":"uuid.UUID","Kind":17,"PkgPath":"github.com/google/uuid","Methods":{"ClockSequence":{"In":[],"Out":[{"Name":"int","Ident":"int","Kind":2,"PkgPath":"","Methods":null}]},"Domain":{"In":[],"Out":[{"Name":"Domain","Ident":"uuid.Domain","Kind":8,"PkgPath":"github.com/google/uuid","Methods":null}]},"ID":{"In":[],"Out":[{"Name":"uint32","Ident":"uint32","Kind":10,"PkgPath":"","Methods":null}]},"MarshalBinary":{"In":[],"Out":[{"Name":"","Ident":"[]uint8","Kind":23,"PkgPath":"","Methods":null},{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"MarshalText":{"In":[],"Out":[{"Name":"","Ident":"[]uint8","Kind":23,"PkgPath":"","Methods":null},{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"NodeID":{"In":[],"Out":[{"Name":"","Ident":"[]uint8","Kind":23,"PkgPath":"","Methods":null}]},"Scan":{"In":[{"Name":"","Ident":"interface {}","Kind":20,"PkgPath":"","Methods":null}],"Out":[{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"String":{"In":[],"Out":[{"Name":"string","Ident":"string","Kind":24,"PkgPath":"","Methods":null}]},"Time":{"In":[],"Out":[{"Name":"Time","Ident":"uuid.Time","Kind":6,"PkgPath":"github.com/google/uuid","Methods":null}]},"URN":{"In":[],"Out":[{"Name":"string","Ident":"string","Kind":24,"PkgPath":"","Methods":null}]},"UnmarshalBinary":{"In":[{"Name":"","Ident":"[]uint8","Kind":23,"PkgPath":"","Methods":null}],"Out":[{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"UnmarshalText":{"In":[{"Name":"","Ident":"[]uint8","Kind":23,"PkgPath":"","Methods":null}],"Out":[{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"Value":{"In":[],"Out":[{"Name":"Value","Ident":"driver.Value","Kind":20,"PkgPath":"database/sql/driver","Methods":null},{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"Variant":{"In":[],"Out":[{"Name":"Variant","Ident":"uuid.Variant","Kind":8,"PkgPath":"github.com/google/uuid","Methods":null}]},"Version":{"In":[],"Out":[{"Name":"Version","Ident":"uuid.Version","Kind":8,"PkgPath":"github.com/google/uuid","Methods":null}]}}}},"immutable":true,"position":{"Index":2,"MixedIn":false,"MixinIndex":0},"comment":"标识符"},{"name":"type","type":{"Type":6,"Ident":"console.Type","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"enums":[{"N":"exchange","V":"exchange"},{"N":"operate","V":"operate"},{"N":"cabinet","V":"cabinet"}],"position":{"Index":3,"MixedIn":false,"MixinIndex":0},"comment":"日志类别 exchange:换电控制 operate:手动操作 cabinet:电柜日志"},{"name":"user_id","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":4,"MixedIn":false,"MixinIndex":0},"comment":"用户ID"},{"name":"user_type","type":{"Type":8,"Ident":"model.UserType","PkgPath":"github.com/auroraride/adapter/model","PkgName":"model","Nillable":false,"RType":{"Name":"UserType","Ident":"model.UserType","Kind":24,"PkgPath":"github.com/auroraride/adapter/model","Methods":{"Scan":{"In":[{"Name":"","Ident":"interface {}","Kind":20,"PkgPath":"","Methods":null}],"Out":[{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"String":{"In":[],"Out":[{"Name":"string","Ident":"string","Kind":24,"PkgPath":"","Methods":null}]},"Value":{"In":[],"Out":[{"Name":"Value","Ident":"driver.Value","Kind":20,"PkgPath":"database/sql/driver","Methods":null},{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]}}}},"position":{"Index":5,"MixedIn":false,"MixinIndex":0},"schema_type":{"postgres":"varchar"},"comment":"用户类别"},{"name":"step","type":{"Type":8,"Ident":"model.ExchangeStep","PkgPath":"github.com/auroraride/adapter/model","PkgName":"model","Nillable":false,"RType":{"Name":"ExchangeStep","Ident":"model.ExchangeStep","Kind":8,"PkgPath":"github.com/auroraride/adapter/model","Methods":{"Index":{"In":[],"Out":[{"Name":"int","Ident":"int","Kind":2,"PkgPath":"","Methods":null}]},"Scan":{"In":[{"Name":"","Ident":"interface {}","Kind":20,"PkgPath":"","Methods":null}],"Out":[{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"String":{"In":[],"Out":[{"Name":"string","Ident":"string","Kind":24,"PkgPath":"","Methods":null}]},"Value":{"In":[],"Out":[{"Name":"Value","Ident":"driver.Value","Kind":20,"PkgPath":"database/sql/driver","Methods":null},{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]}}}},"nillable":true,"optional":true,"position":{"Index":6,"MixedIn":false,"MixinIndex":0},"schema_type":{"postgres":"smallint"},"comment":"换电步骤"},{"name":"status","type":{"Type":6,"Ident":"console.Status","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"enums":[{"N":"invalid","V":"invalid"},{"N":"pending","V":"pending"},{"N":"running","V":"running"},{"N":"success","V":"success"},{"N":"failed","V":"failed"}],"position":{"Index":7,"MixedIn":false,"MixinIndex":0},"comment":"状态 invalid:无效 pending:未开始 running:执行中 success:成功 failed:失败"},{"name":"before_bin","type":{"Type":3,"Ident":"*model.BinInfo","PkgPath":"github.com/auroraride/adapter/model","PkgName":"model","Nillable":true,"RType":{"Name":"BinInfo","Ident":"model.BinInfo","Kind":22,"PkgPath":"github.com/auroraride/adapter/model","Methods":{}}},"optional":true,"position":{"Index":8,"MixedIn":false,"MixinIndex":0},"comment":"变化前仓位信息"},{"name":"after_bin","type":{"Type":3,"Ident":"*model.BinInfo","PkgPath":"github.com/auroraride/adapter/model","PkgName":"model","Nillable":true,"RType":{"Name":"BinInfo","Ident":"model.BinInfo","Kind":22,"PkgPath":"github.com/auroraride/adapter/model","Methods":{}}},"optional":true,"position":{"Index":9,"MixedIn":false,"MixinIndex":0},"comment":"变化后仓位信息"},{"name":"message","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"nillable":true,"optional":true,"position":{"Index":10,"MixedIn":false,"MixinIndex":0},"comment":"消息"},{"name":"startAt","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"nillable":true,"optional":true,"position":{"Index":11,"MixedIn":false,"MixinIndex":0},"comment":"开始时间"},{"name":"stopAt","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"nillable":true,"optional":true,"position":{"Index":12,"MixedIn":false,"MixinIndex":0},"comment":"结束时间"},{"name":"duration","type":{"Type":20,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"nillable":true,"optional":true,"position":{"Index":13,"MixedIn":false,"MixinIndex":0},"comment":"耗时"}],"indexes":[{"fields":["cabinet_id"]},{"fields":["bin_id"]},{"fields":["serial"]},{"fields":["uuid"]},{"fields":["user_id"]},{"fields":["user_type"]},{"fields":["startAt"]},{"fields":["stopAt"]}],"annotations":{"EntSQL":{"table":"console","with_comments":true}}},{"name":"Scan","config":{"Table":""},"edges":[{"name":"cabinet","type":"Cabinet","field":"cabinet_id","unique":true,"required":true}],"fields":[{"name":"created_at","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"default":true,"default_kind":19,"immutable":true,"position":{"Index":0,"MixedIn":true,"MixinIndex":0}},{"name":"updated_at","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"default":true,"default_kind":19,"update_default":true,"position":{"Index":1,"MixedIn":true,"MixinIndex":0}},{"name":"cabinet_id","type":{"Type":18,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":0,"MixedIn":true,"MixinIndex":1}},{"name":"id","type":{"Type":4,"Ident":"uuid.UUID","PkgPath":"github.com/google/uuid","PkgName":"uuid","Nillable":false,"RType":{"Name":"UUID","Ident":"uuid.UUID","Kind":17,"PkgPath":"github.com/google/uuid","Methods":{"ClockSequence":{"In":[],"Out":[{"Name":"int","Ident":"int","Kind":2,"PkgPath":"","Methods":null}]},"Domain":{"In":[],"Out":[{"Name":"Domain","Ident":"uuid.Domain","Kind":8,"PkgPath":"github.com/google/uuid","Methods":null}]},"ID":{"In":[],"Out":[{"Name":"uint32","Ident":"uint32","Kind":10,"PkgPath":"","Methods":null}]},"MarshalBinary":{"In":[],"Out":[{"Name":"","Ident":"[]uint8","Kind":23,"PkgPath":"","Methods":null},{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"MarshalText":{"In":[],"Out":[{"Name":"","Ident":"[]uint8","Kind":23,"PkgPath":"","Methods":null},{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"NodeID":{"In":[],"Out":[{"Name":"","Ident":"[]uint8","Kind":23,"PkgPath":"","Methods":null}]},"Scan":{"In":[{"Name":"","Ident":"interface {}","Kind":20,"PkgPath":"","Methods":null}],"Out":[{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"String":{"In":[],"Out":[{"Name":"string","Ident":"string","Kind":24,"PkgPath":"","Methods":null}]},"Time":{"In":[],"Out":[{"Name":"Time","Ident":"uuid.Time","Kind":6,"PkgPath":"github.com/google/uuid","Methods":null}]},"URN":{"In":[],"Out":[{"Name":"string","Ident":"string","Kind":24,"PkgPath":"","Methods":null}]},"UnmarshalBinary":{"In":[{"Name":"","Ident":"[]uint8","Kind":23,"PkgPath":"","Methods":null}],"Out":[{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"UnmarshalText":{"In":[{"Name":"","Ident":"[]uint8","Kind":23,"PkgPath":"","Methods":null}],"Out":[{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"Value":{"In":[],"Out":[{"Name":"Value","Ident":"driver.Value","Kind":20,"PkgPath":"database/sql/driver","Methods":null},{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"Variant":{"In":[],"Out":[{"Name":"Variant","Ident":"uuid.Variant","Kind":8,"PkgPath":"github.com/google/uuid","Methods":null}]},"Version":{"In":[],"Out":[{"Name":"Version","Ident":"uuid.Version","Kind":8,"PkgPath":"github.com/google/uuid","Methods":null}]}}}},"default":true,"default_kind":19,"position":{"Index":0,"MixedIn":false,"MixinIndex":0}},{"name":"efficient","type":{"Type":1,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"default":true,"default_value":true,"default_kind":1,"position":{"Index":1,"MixedIn":false,"MixinIndex":0},"comment":"是否有效"},{"name":"user_id","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":2,"MixedIn":false,"MixinIndex":0},"comment":"用户ID"},{"name":"user_type","type":{"Type":8,"Ident":"model.UserType","PkgPath":"github.com/auroraride/adapter/model","PkgName":"model","Nillable":false,"RType":{"Name":"UserType","Ident":"model.UserType","Kind":24,"PkgPath":"github.com/auroraride/adapter/model","Methods":{"Scan":{"In":[{"Name":"","Ident":"interface {}","Kind":20,"PkgPath":"","Methods":null}],"Out":[{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"String":{"In":[],"Out":[{"Name":"string","Ident":"string","Kind":24,"PkgPath":"","Methods":null}]},"Value":{"In":[],"Out":[{"Name":"Value","Ident":"driver.Value","Kind":20,"PkgPath":"database/sql/driver","Methods":null},{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]}}}},"position":{"Index":3,"MixedIn":false,"MixinIndex":0},"schema_type":{"postgres":"varchar"},"comment":"用户类别"},{"name":"serial","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":4,"MixedIn":false,"MixinIndex":0},"comment":"电柜编号"},{"name":"data","type":{"Type":3,"Ident":"*model.ExchangeUsableResponse","PkgPath":"github.com/auroraride/adapter/model","PkgName":"model","Nillable":true,"RType":{"Name":"ExchangeUsableResponse","Ident":"model.ExchangeUsableResponse","Kind":22,"PkgPath":"github.com/auroraride/adapter/model","Methods":{}}},"optional":true,"position":{"Index":5,"MixedIn":false,"MixinIndex":0},"comment":"换电信息"}],"indexes":[{"fields":["created_at"]},{"fields":["cabinet_id"]},{"fields":["user_id"]},{"fields":["user_type"]},{"fields":["serial"]}],"annotations":{"EntSQL":{"table":"scan","with_comments":true}}}],"Features":["sql/modifier","sql/upsert","privacy","entql","sql/execquery","schema/snapshot"]}`
