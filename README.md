# URL Shortener - Week 1 Go Learning Project

## Giới thiệu
URL Shortener đơn giản được viết bằng Go - rút gọn liên kết dài thành liên kết ngắn và theo dõi số lượt truy cập.

## Chức năng
- 🔗 Rút gọn URL dài thành URL ngắn  
- 📊 Theo dõi số lượt click
- 🔄 Redirect từ URL ngắn về URL gốc
- 📋 Xem danh sách và thống kê tất cả URL

## Cách chạy

### 1. Cài đặt dependencies
```bash
go mod download
```

### 2. Chạy ứng dụng
```bash
go run main.go
```

### 3. Mở trình duyệt
Truy cập: `http://localhost:8080`

## Sử dụng

### Giao diện web
- Mở `http://localhost:8080` trong trình duyệt
- Nhập URL cần rút gọn và click "Rút gọn"
- Xem danh sách URL đã tạo bên dưới
- Click "Copy" để copy URL ngắn
- Click "Mở" để test redirect
