1. Data racing
    - Xảy ra khi có nhiều luồng cùng truy cập và thao tác trên cùng một resource.
2. Atomicity
    - Tính nguyên tử của resource.
    - Nếu sử dụng Lock và Unlock dể synchronize access tới resource trong Mutex thì không thực sự xử lý được data races do vẫn không thể biết chính xác thứ tự luồng nào truy cập vào resource trước -> k đảm bảo logic xử lý.
    - Sử dụng lock cũng gây slow chương trình
    - 2 câu hỏi đặt ra: Liệu các phần code synchronize có bị truy cập liên tục và phạm vi phần code synchronize như thế nào ?
3. Deadlock, Livelocks, Starvation
4. sync package
   - WaitGroup: 
     - Sử dụng để đợi một tập hợp các concurrency routine mà không quan tâm đến kết quả của concurrency data hoặc có cách khác để lấy được
     - Call Add() trước khi tạo routine để tránh Add() không được call do code có thể chạy tới Wait() trước khi nhảy vào routine
     - Call Done() with defer
   - Mutex and RWMutex:
     - Call Lock, Unlock để đồng bộ hoá truy cập tới cùng một resource, tránh race condition.
     - Call Unlock with defer
     - Sử dụng RWMutex.RLock và RUnlock cho các tác vụ đọc; Lock và Unlock cho tác vụ ghi do các tác vụ đọc k cần đợi nhau để sử dụng khoá. Chỉ khi tác vụ giữ khoá là tác vụ ghi thì mới phải đợi.