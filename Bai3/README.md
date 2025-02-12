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
# endpoint thứ 1 người dùng bảo con AI tạo hội thoại ngẫu nhiên http://localhost:8080/api/dialog/process
{
  "prompt": "Tạo một hội thoại tiếng Việt về hỏi đường đến hồ Hoàn Kiếm giữa James và Lan."
}

![alt text](./endpointB3.png)
# endpoint thứ 2 cho người dùng tự nhập hội thoại http://localhost:8080/api/dialog/manual
{
  "content": "James: Chào bạn! Bạn có thể chỉ tôi đường đến hồ Hoàn Kiếm không?\nLan: Chào bạn! Bạn đang ở đâu?\nJames: Tôi đang ở phố Tràng Tiền.\nLan: Từ đây, bạn đi thẳng, rẽ phải vào đường Đinh Tiên Hoàng, sẽ thấy hồ.\nJames: Cảm ơn bạn nhiều!\nLan: Không có gì, chúc bạn đi vui!"
}
# kết quả 
![alt text](./endpointB3thu2.png.png)
