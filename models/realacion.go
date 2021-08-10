package models

/*Relacion entre 2 usuarios*/
type Relacion struct {
	UsuarioID         string `bson:"usuarioid" json:"usuarioId"`
	UsuarioRelacionID string `bson:"usuariorelacionid" json:"usuarioRelacionId"`
}

/*RespuspuestaConsultaRelacion tiene el true o false que se obtiene de consultar la ralacion entre 2 usuarios */
type RespuspuestaConsultaRelacion struct {
	Status bool `json:"status"`
}
