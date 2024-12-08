# Xây dựng api cho ứng dụng đặt vé xem phim bằng Golang (GORM)
## Mô tả dự án
- Xây dựng hệ thống back-end cho ứng dụng đặt vé xem phim
- Sử dụng: Golang Với Gin, Gorm, PostgreSQL

- Thiết kế Database 
    [Link Dbdiagram](https://dbdiagram.io/d/5de9c01cedf08a25543ec5c0)
    <img src="https://i.imgur.com/ScziiJD.png">

## Các chức năng chính của hệ thống
- Lấy dữ liệu thông tin phim đang chiếu
    > `http://localhost/movies/now` 
    <img src="https://i.imgur.com/tHBnbOz.jpg">
- Lấy dữ liệu thông tin phim sắp chiếu
    > `http://localhost/movies/future` 
    <img src="https://i.imgur.com/uwcUX9M.jpg">
- Lấy dữ liệu thông tin các suất chiếu đặc biệt
- Lấy thông tin các suất chiếu
    > `localhost/schedule/movie_id/date`
                        
   - Trong đó `movie_id` là id của bộ phim, `date` là ngày (định dạng YYYY-mm-dd) có suất chiếu của bộ phim `movie_id`. 
  
  <img src="https://i.imgur.com/hAx2KkZ.jpg">
- Thực hiện đặt vé
- Đăng kí, đăng nhập 

## Thực hiện cài đặt
- Thực hiện lệnh 
    > `git clone https://github.com/Biiisme/Cinema.git`

## Các lỗi gặp phải nếu có

### Không kết nối được đến PostgresSQL
- Cách fix lỗi: sử dụng `.env` để cấu hình lại host 
## Các việc trong tuần tới
- Tiến hành ghép api với front-end
 