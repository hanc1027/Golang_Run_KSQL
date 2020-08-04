## Golang_Run_ksql

### 需在go環境下執行
### setup
1) 請至[Golang官網](https://golang.org/dl/)下載對應的OS版本
2) 設定環境變數
    1) 將/usr/local/go/bin加入PATH環境變數，將以下的設定寫入*.bashrc*。  
    若有使用iTerm2，則寫入*.zshrc*

        ```shell
            export GOPATH=$HOME/go
            export PATH=$HOME/bin:$GOPATH/bin:$PATH
        ```

    2) 執行 `$ source .bashrc` => 讓設定值重新跑

    3) 可再執行`$ echo $PATH` => 檢查環境變數是否有設定成功
3) 下載、安裝完成後，於Mac或Linux系統中，找到 **/Users/[您的使用者名稱]/go/** 的路徑。  
Windows系統，應可在**C:/go**下找到。
4) 將此專案移至`/Users/[您的使用者名稱]/go/src`的資料下，再執行`$ go run .`，即可啟動KSQL Statement。
