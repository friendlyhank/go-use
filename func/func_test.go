package _func

/**
1.func可以用于那些地方

	(1)用于type定义
	type ReadDirOption func(*ReadDirOp)

	(2)在struct
	type File struct{
		ReadDirOption func(*ReadDirOp)
	}

	(3)return
	func WithExt(ext string)ReadDirOption{
		return func(op *ReadDirOp){
			op.ext = ext
		}
	}

 */

