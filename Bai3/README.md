<!-- đây chỉ là phần code BE chứa endpoint -->
# chạy lệnh này để khởi tạo
<!-- go mod init Bai3
go get github.com/kataras/iris/v12
go get github.com/lib/pq -->
go get github.com/gin-gonic/gin
go get 	"github.com/spf13/viper"

# tạo database trong db ten Bai3
CREATE TABLE dialog (
    id BIGSERIAL PRIMARY KEY,
    lang VARCHAR(2) NOT NULL,
    content TEXT NOT NULL
);

CREATE TABLE word (
    id BIGSERIAL PRIMARY KEY,
    lang VARCHAR(2) NOT NULL,
    content TEXT NOT NULL,
    translate TEXT NOT NULL
);

CREATE TABLE word_dialog (
    dialog_id BIGINT REFERENCES dialog(id) ON DELETE CASCADE,
    word_id BIGINT REFERENCES word(id) ON DELETE CASCADE,
    PRIMARY KEY (dialog_id, word_id)
);
# hãy sửa file config.yaml cho đúng với db và key của bạn
# đây chính là kết quả endpoint chạy được ở bài 3 ta sẽ tách riêng FE và BE ra
![alt text](./endpointB3.png)
