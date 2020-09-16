# Assignment #2: gRPC and REST API Benchmarking
## Member

1. Natcha Manasuntorn 6030177021
2. Suchut Sapsathien 6030609921
3. Karnkitti Kittikamron 6031006621
4. Yanika Dontong 6031010021
5. Natthanon Manop 6031013021
## Things to be delivered:
1. Graphs showing the benchmarking results with the explanation of your experimental settings. </br>
* Single client with a small call to insert a book item, a bigger call to insert a list of multiple book items. </br>
![](https://i.imgur.com/3j09doy.png)
* Multiple clients with different kind of calls </br>
![](https://i.imgur.com/grr5D4y.png)
* Vary the number of concurrent calls from 1 to 4096 calls. </br>
![](https://i.imgur.com/oHVokAu.png)
---

2. Discussion of the results why one method is better the other in which scenarios.</br>
**จากกราฟที่ 1** แสดงให้เห็นว่า gRPC ตอบสนองเร็วกว่า RestAPI ในการเรียกคำสั่งที่มีขนาดใหญ่ 
**จากกราฟที่ 2** แสดงให้เห็นว่า gRPC ตอบสนองเร็วกว่า RestAPI เมื่อมีการเรียกใช้งานคำสั่งต่างๆจากหลายๆclient 
**จากกราฟที่ 3** แสดงให้เห็นว่าเมื่อมี client เรียกใช้คำสั่งหลายๆclientพร้อมกันแล้วการตอบสนองของทั้งคู่จะมีค่าที่ใกล้เคียงกัน </br>
จะเห็นได้ว่า gRPC มีประสิทธิภาพที่ดีกว่า Rest API ในหลายกรณี โดยที่ gRPC นั้นใช้มาตรฐาน HTTP/2 ซึ่งรองรับการเรียก request พร้อมกันทำให้ ทำงานได้เร็วกว่า Rest API ซึ่งใช้มาตรฐาน HTTP/1.1 ที่ไม่รองรับ การเรียก request พร้อมกัน แต่เนื่องจากประสิทธิภาพของหน่วยประมวลผลทำให้ได้ผลลัพธ์ออกมาใกล้เคียงกัน นอกจากนี้ gRPC ใช้ข้อมูลในการทำงานในรูปแบบของ binary Protobuf ซึ่งมีขนาดที่เล็กกว่า ข้อมูลรูปแบบ JSON ของ Rest API จึงทำให้มีประสิทธิภาพการทำงานที่ดีกว่าและเร็วกว่า

---

3. Comparison of the gRPC and REST API from the aspects of language neutral, ease of use, and performance. 
    - Language neutral
        - Rest API: เนื่องจาก Rest API สามารถใช้งานข้ามภาษาได้และมีการใช้งานกันมาเป็นระยะเวลานาน จึงทำ Rest API มีภาษาที่รับรองค่อนข้างมาก
        - gRPC: gRPC สามารถทำงานข้ามภาษาได้เหมือนกันกับ Rest API และมีการรับรองจากภาษาหลักที่นิยมใช้กัน เช่น C++, Java, Python, Go, Ruby, C#, Node.js, Android Java, Objective-C, PHP, Dart
    - Ease of use
        - Rest API: มี http method ให้เรียกใช้ได้แก่ GET/POST/PUT/DELETE และตอนส่งrequestจะต้องมีการกำหนดpathย่อย,headerและ body 
        - gRPC: สามารถกำหนดmethodที่เรียกใช้ได้เอง และตอนส่งrequestไม่ต้องมีการกำหนดpathย่อย,headerและ body ทำให้ใช้งานได้ง่ายยิ่งขึ้น
    - Performance
        - Rest api: จะมีประสิทธิภาพที่ต่ำกว่า(ช้ากว่า) gRPC อย่างเห็นได้ชัด หากมีการเรียกคำสั่งขนาดใหญ่ เนื่องจากใช้ HTTP/1.1.
        - gRPC: จะมีประสิทธิภาพที่สูงกว่า(เร็วกว่า) Rest อย่างเห็นได้ชัด หากมีการเรียกคำสั่งขนาดใหญ่ เนื่องจากใช้ HTTP/2.
---

4. Does your results comply with the results in https://medium.com/@bimeshde/grpc-vs-rest-performance-simplifiedfd35d01bbd4? How? </br>
ใช่ ผลลัพธ์มีความเป็นไปในทิศทางเดียวกันกับ medium โดยที่ในส่วนของประสิทธิภาพ gRPC จะมีประสิทธิภาพที่ดีกว่า Rest API และในส่วนของ ease of use จะขึ้นกับว่านักพัฒนา มีความสามารถหรือมีความต้องการพัฒนาบนภาษาไหนมากกว่ากัน

---
