## Работаем с гитом в памяти. Делаем пулл прямиком в Zip архив
- При отключенном параметре GO111MODULE = off, выполнить тестовый пример из директории vgo_disabled  
Используем модифицированный GOPATH, что-бы не настраивать go modules

1) import пакета (go get github.com/go-git/go-billy)
2) В GOPATH находим пакет go-git (GOPATH\src\github.com\go-git\go-billy), создаем папку v5
3) Копируем туда содержимое из этой папки  
Все работает.  
<img src="modGopath.gif" width="600" />  
Или можно ипользовать мой https://yadi.sk/d/yUaV7m5jjvCfZg  


- При включенном параметре GO111MODULE = on, можно просто выролнить тестовый пример из директории vgo_support  

