package models

var JsonDefaultModels = [][]byte{
	[]byte(`
    {
        "model_name": "5 STEPS MODEL",
        "model_nodes": ["S1","S2","S3","S4","S5"],
        "model_edges": [{"S1":"S2"},{"S1":"S3"},{"S1":"S4"},{"S2":"S5"},{"S3":"S5"},{"S4":"S5"}]
    }`),
}
