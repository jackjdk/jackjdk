package main

func main() {
	/*
		拷贝的对象：深拷贝拷贝的是数据本身，而浅拷贝拷贝的是数据地址。因此，深拷贝会创建一个新的对象，该对象与原对象不共享内存，新对象值修改时不会影响
		原对象值。而浅拷贝则是只复制指向的对象的指针，新对象和老对象指向的内存地址是一样的，新对象值修改时老对象也会变化。

		内存地址：由于深拷贝和浅拷贝拷贝的对象不同，因此在内存中的表现也有所不同。深拷贝在内存中开辟一个新的内存地址，而浅拷贝则是复制指向的对象的指针，
		因此释放内存地址时，深拷贝可以分别释放，而浅拷贝则会同时释放。

		数据类型：深拷贝和浅拷贝的特性并不仅限于特定的数据类型，而是取决于数据类型的实现。也就是说，一个数据类型如果实现了深拷贝的方法，那么该数据类型
		的实例在拷贝时就会进行深拷贝。如果没有实现深拷贝的方法，那么默认情况下进行的就是浅拷贝。
	*/
}
