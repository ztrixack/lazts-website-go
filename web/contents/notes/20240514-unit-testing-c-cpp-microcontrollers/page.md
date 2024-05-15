---
Title: Test
Slug: 20240514-unit-testing-c-cpp-microcontrollers
Excerpt: Add YAML metadata to the document
FeaturedImage: https://picsum.photos/1024/768?random=12
FeaturedImageAlt: Sample Images
PublishedAt: 2024-05-14
Published: true
Tags:
  - embedded
  - test
  - C/C++
---

# Unit Testing C/C++ for Microcontrollers

Unit Testing เป็นแนวทางหนึ่งใน software development เพื่อให้มั่นใจว่าส่วนประกอบแต่ละส่วนของโค้ดจะทำงานได้ตามที่ต้องการ และสามารถ maintain โค้ดชุดนั้นได้ต่อเนื่องยาวนาน คนใหม่ๆ ที่เข้ามาดูแลต่อก็จะมั่นใจว่าจะไม่ทำระบบพัง ซึ่งก็เป็นแนวทางที่ปรมาจารย์ทุกท่านแนะนำแกมบังคับให้ทำตาม

แต่ถึงแม้ว่ามจะดีแค่ไหน การเขียน unit testing ก็ยากและขัดใจเหลือเกิน โดยเฉพาะเมื่อต้องทำงานร่วมกับ hardware

ในการเขียน firmware บน microcontroller ส่วนใหญ่จะเน้นไปที่การควบคุมการทำงานของ hardware หรืออุปกรณ์ต่างๆ ไม่ว่าจะเป็นการอ่าน/เขียน register การตอบสนองต่อ event หรือ interrupt จากภายนอก ซึ่งนักพัฒนาอย่างเราๆ ก็จะต้องทำการเทสและดีบักควบคู่กันไปกับอุปกรณ์ต่างๆ ด้วยอยู่แล้ว เมื่อมองในแง่ของการเทส มันก็คือ end-to-end testing ร่วมกับ hardware ตลอดเวลา การเขียน unit testing มาทดสอบอีกครั้งหลังจากทำ end-to-end testing แล้วจึงเป็นอะไรที่ไม่คุ้มค่าเหนื่อยเลย จึงไม่แปลกที่นักพัฒนาจะเริ่มต้นจากการหาวิธีที่มันจะทำงานได้ตามต้องการ ปรับจูนทุกอย่างให้ได้ก่อน เพราะแม้ end-to-end testing จะมี cost สูงแต่สิ่งที่ได้รับคือความสบายใจ

ดังนั้นการที่จะได้เห็น unit testing บน microcontroller จึงเป็นอะไรที่แรร์สุดๆ



แต่ถ้ามองในอีกมุมหนึ่งแล้ว
โค้ดบน microcontroller ก็ไม่ได้ต่างจากระบบอื่นๆ เลย

ถ้าเราแยกส่วนของ Hardware การเขียน Register, Interrupt ฯลฯ ออกมาให้อยู่ในรูปของ Module/Service หรืออะไรก็ตามแต่

ดังนั้น สิ่งที่สำคัญอยู่ที่การออกแบบโครงสร้างของ firmware ถ้าออกแบบมาอย่างดีและสามารถแยกชั้นฮาร์ดแวร์และแอปพลิเคชันได้อย่างชัดเจน ก็จะทำให้การเขียน unit test เป็นเรื่องง่ายขึ้น ส่วนของ hardware ก็ใช้ logic analyzer, osciloscope, หรืออุปกรณ์อื่นๆ ส่วนที่เหลือก็จะเป็นเพียงโค้ด C/C++ เพียวๆ ที่สามารถ test บนระบบได้

## Why Unit Testing is Essential for Microcontroller Development

หา bug ได้เร็ว: ช่วยให้ค้นพบข้อผิดพลาดได้เร็วขึ้น ลดค่าใช้จ่ายและความพยายามในการแก้ไข
ปรับปรุงคุณภาพโค้ด: ส่งเสริมการเขียนโค้ดและการออกแบบที่ดีขึ้น
ป้องกันการเกิดบั๊กซ้ำ: ทำให้มั่นใจว่าการเปลี่ยนแปลงใหม่จะไม่ทำให้ฟังก์ชันการทำงานที่มีอยู่เสียหาย
เอกสารประกอบ: ให้ตัวอย่างที่ชัดเจนของฟังก์ชันและพฤติกรรมของโมดูลที่คาดหวัง
เพิ่มความมั่นใจในโค้ด: มั่นใจว่าโค้ดทำงานตามที่ตั้งใจไว้ แม้จะมีการเปลี่ยนแปลง
สนับสนุนการ: รวมเข้ากับกระบวนการ CI สำหรับการทดสอบและการตรวจสอบอัตโนมัติ

## Why Unit Test Microcontroller Code?
Microcontrollers are used in critical applications where reliability is paramount. Unit testing helps to:

Identify and fix bugs early in the development process.
Ensure that changes or additions to the code do not break existing functionality.
Facilitate code refactoring and optimization with confidence.

## Unit Test Framework Comparison

| Framework    | Language | Lightweight | Mocking Framework   | Embedded | Comments                   |
| ------------ | -------- | ----------- | ------------------- | -------- | -------------------------- |
| [CUnit]      | C        | Yes         | No                  | Limited  | good for basic testing     |
| [CMock]      | C        | Yes         | Yes                 | Yes      | enhances isolation         |
| [Unity]      | C        | Yes         | Yes (CMock)         | Yes      | minimal overhead           |
| [Ceedling]   | C        | Yes         | Yes (Unity & CMock) | Yes      | streamlines workflow       |
| [GoogleTest] | C++      | No          | Yes (GMock)         | Limited  | powerful, well-document    |
| [CppUTest]   | C/C++    | Yes         | Yes (CppUMock)      | Yes      | flexible, active community |
| [Check]      | C        | Yes         | No                  | Limited  | easy to use, good for CI   |


## Importance of Unit Testing in Microcontroller Development

Unit testing in microcontroller development helps to:
- Identify and fix bugs early in the development process.
- Ensure that individual modules of code work correctly before integration.
- Facilitate maintenance and refactoring of code.
- Improve code reliability and robustness.

## Implementing Unit Testing in Microcontroller Projects

### 1. **Setting Up the Environment**
   - Choose a suitable unit testing framework based on project requirements and constraints.
   - Integrate the framework with the build system (e.g., Ceedling for a complete TDD setup).

### 2. **Writing Tests**
   - Begin with writing test cases for individual functions.
   - Use assertions to validate expected outcomes.
   - Ensure tests are isolated and do not depend on external states.

### 3. **Mocking and Stubbing**
   - Use tools like CMock to create mocks and stubs for dependencies.
   - This allows for testing functions in isolation, especially useful when dealing with hardware-specific code.

### 4. **Continuous Integration**
   - Integrate unit tests with a continuous integration (CI) system.
   - Automate test execution to catch regressions early.

### 5. **Test-Driven Development (TDD)**
   - Adopt TDD practices by writing tests before the actual code.
   - This ensures that code is always written to satisfy the test requirements, leading to better design and fewer defects.

## Example Workflow with Ceedling

1. **Initialize the Project**
   ```sh
   ceedling new project_name
   cd project_name
   ```

2. **Write a Test Case**
   ```c
   // test/test_example.c
   #include "unity.h"
   #include "example.h"

   void test_addition(void) {
       TEST_ASSERT_EQUAL(4, add(2, 2));
   }
   ```

3. **Run Tests**
   ```sh
   ceedling test:all
   ```

## When to Use Mocking

Mocking is essential when:
- The code interacts with hardware peripherals.
- Dependencies are complex or not yet implemented.
- Testing interactions with external systems.

## Conclusion

Unit testing in microcontroller development is crucial for building reliable and maintainable software. By leveraging frameworks like Unity, CMock, and Ceedling, developers can implement effective unit tests and adopt best practices such as TDD and continuous integration. These practices ensure that embedded software is robust, with fewer bugs and easier maintenance.

By following the guidelines and utilizing the right tools, you can enhance the quality of your microcontroller projects and deliver reliable embedded systems.

## Reference
[How To Implement C Programming Unit Testing Successfully](https://marketsplash.com/c-programming-unit-testing/)
[What is the Best Unit Testing Framework in C for You?](https://moderncprogramming.com/what-is-the-best-unit-testing-framework-in-c-for-you/)
[Modern unit testing in C with TDD and Ceedling](https://www.embedded.com/modern-unit-testing-in-c-with-tdd-and-ceedling/)
[When to Mock Unit Testing C/C++ Code](https://www.parasoft.com/blog/unit-testing-c-code-when-to-mock/)
[Test-driven development and unit testing with examples in C++](https://alexott.net/en/cpp/CppTestingIntro.html)
[List of unit testing frameworks](https://en.wikipedia.org/wiki/List_of_unit_testing_frameworks#C++)

[CUnit]: http://cunit.sourceforge.net/
[Unity]: http://www.throwtheswitch.org/unity
[CMock]: https://github.com/ThrowTheSwitch/CMock
[Ceedling]: https://github.com/ThrowTheSwitch/Ceedling
[GoogleTest]: https://github.com/google/googletest
[CppUTest]: http://cpputest.github.io/
[Check]: https://libcheck.github.io/check/
