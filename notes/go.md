### Go 语法补充

- omitempty 关键字 
  - omitempty是省略的意思 
  - json中字段若有omitempty标记，则这个字段为空时，json序列化为string时不会包含该字段 
  - json中字段若没有omitempty标记，则这个字段为空时，json序列化为string时会包含该字段