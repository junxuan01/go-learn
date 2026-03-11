// ========== JavaScript 的 Class 方式 ==========

// 基类 Person
class Person {
  constructor(name, age) {
    this.name = name;
    this.age = age;
    this.birthDate = new Date();
    this.birthDate.setFullYear(this.birthDate.getFullYear() - age);
  }

  // 方法
  introduce() {
    return `我叫 ${this.name}，今年 ${this.age} 岁`;
  }

  getAge() {
    return this.age;
  }
}

// Student 继承 Person
class Student extends Person {
  constructor(name, age, studentID, grade) {
    super(name, age); // 调用父类构造函数
    this.studentID = studentID;
    this.grade = grade;
    this.scores = new Map();
    this.enrollmentDate = new Date();
  }

  // Student 自己的方法
  study(subject) {
    console.log(`${this.name} 正在学习 ${subject}`);
  }

  addScore(subject, score) {
    this.scores.set(subject, score);
  }

  getAverageScore() {
    if (this.scores.size === 0) return 0;
    let total = 0;
    for (const score of this.scores.values()) {
      total += score;
    }
    return total / this.scores.size;
  }

  // 重写父类方法
  introduce() {
    return `我叫 ${this.name}，今年 ${this.age} 岁，学号 ${this.studentID}，年级 ${this.grade}`;
  }

  // 可以调用父类方法
  introduceAsParent() {
    return super.introduce();
  }
}

// GraduateStudent 继承 Student（多层继承）
class GraduateStudent extends Student {
  constructor(name, age, studentID, researchTopic, advisor) {
    super(name, age, studentID, 5); // 研究生默认5年级
    this.researchTopic = researchTopic;
    this.advisor = advisor;
  }

  doResearch() {
    console.log(`${this.name} 正在研究 ${this.researchTopic}，导师是 ${this.advisor}`);
  }

  // 研究生的自我介绍进一步重写
  introduce() {
    return `${super.introduce()}，研究方向: ${this.researchTopic}`;
  }
}

// ClassRoom 类（组合方式）
class ClassRoom {
  constructor(roomNumber, teacher) {
    this.roomNumber = roomNumber;
    this.students = [];
    this.teacher = teacher; // 组合一个 Person 对象
  }

  addStudent(student) {
    this.students.push(student);
  }

  listStudents() {
    console.log(`教室 ${this.roomNumber} 的学生列表：`);
    this.students.forEach((student, index) => {
      console.log(`  ${index + 1}. ${student.name} (学号: ${student.studentID})`);
    });
  }
}

// ========== 使用示例 ==========

console.log("========== JavaScript 的 Class 方式示例 ==========\n");

// 1. 创建 Person 实例
const person = new Person("陈老师", 35);
console.log(person.introduce());
console.log();

// 2. 创建 Student 实例
const student1 = new Student("张三", 18, "S001", 2);
const student2 = new Student("李四", 19, "S002", 3);

console.log(student1.introduce());
console.log(student1.introduceAsParent()); // 调用父类方法
console.log();

// 3. 添加成绩
student2.addScore("数学", 95.5);
student2.addScore("英语", 88.0);
student2.addScore("JavaScript", 99.0);
student2.study("算法");
console.log(`${student2.name} 的平均分: ${student2.getAverageScore().toFixed(2)}\n`);

// 4. 研究生示例
const grad = new GraduateStudent("王五", 24, "G001", "分布式系统", "张教授");
console.log(grad.introduce());
grad.doResearch();
grad.study("论文写作"); // 继承自 Student
console.log();

// 5. 教室示例
const classroom = new ClassRoom("101", person);
classroom.addStudent(student1);
classroom.addStudent(student2);
classroom.listStudents();
console.log(`老师: ${classroom.teacher.introduce()}`);

// 6. JavaScript 继承的特点演示
console.log("\n========== JavaScript 继承特点 ==========");

// instanceof 检查
console.log("student1 instanceof Student:", student1 instanceof Student); // true
console.log("student1 instanceof Person:", student1 instanceof Person);   // true
console.log("grad instanceof GraduateStudent:", grad instanceof GraduateStudent); // true
console.log("grad instanceof Student:", grad instanceof Student);         // true
console.log("grad instanceof Person:", grad instanceof Person);           // true

// 原型链
console.log("\n原型链:");
console.log("Student.prototype.__proto__ === Person.prototype:", 
  Student.prototype.__proto__ === Person.prototype); // true

// ========== 对比要点 ==========
console.log("\n========== JavaScript Class vs Go Struct ==========");
console.log(`
JavaScript Class 特点:
1. 使用 extends 关键字实现继承（is-a 关系）
2. 子类是父类的实例 (student instanceof Person === true)
3. 通过原型链查找方法
4. 可以用 super 调用父类方法
5. 单继承链（一个类只能继承一个父类）
6. 继承的是"类型身份"（类型层级）

Go Struct 特点:
1. 使用组合（embedding）实现代码复用（has-a 关系）
2. 不存在类型层级，Student 不是 Person 的子类型
3. 通过字段提升访问嵌入类型的方法
4. 可以通过 s.Person.Method() 显式调用
5. 可以嵌入多个类型（类似多继承，但更清晰）
6. 多态通过接口实现，而不是继承
7. 更灵活，避免深层继承带来的复杂性
`);
