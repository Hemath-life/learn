[What’s SQL Injection]

### What’s SQL Injection ?

- QL Injection is an attack on a web application database server that causes malicious queries to be executed.

- When a web application communicates with a database using input from a user that hasn’t been properly sanitized or validated, An attacker is able to steal, delete or alter private and customer data.

- SQL Injection is one of the oldest yet still one of the most devastating attacks if done right.

### What is a database?

So, we know an SQL attack is performed against a carelessly implemented database. This begs the question, what even is a database?

A database is a way of storing data in an organized manner. Database Management System (DBMS), DBMSs fall into two categories, namely Relational or Non-Relational. To put it simply, Relational databases have some sort of relationship between different tables, tables share information using what’s called a primary key. Whereas, Non-relational don’t use tables or columns. Hence, each row of data can contain different information which can give more flexibility over a relational database.

### Structured Query Language (SQL)

To communicate with databases we use, Structured Query Language (SQL). These SQL queries are better referred to as statements.

The simplest of the commands are select, update, insert and delete data. Some databases servers have their own flavors in syntax and slight changes to how things work. I’ll be using MySQL database. It’s worth noting that SQL syntax is not case-sensitive.

### SQL Injection

This innocent-looking SQL can turn into something devastating when user-provided data gets appended with SQL queries without properly sanitizing it.

You’ve come across an online portal website that your boss uses to keep his company’s data, The entries may be either set to public or private depending on whether they’re ready for public release. The URL for each notice looks something like this

`https://company.com/notice?id=1`

```SQL
SELECT * FROM USERS WHERE id=1 AND role='admin' LIMIT 1;
```

So, it turns out you have direct control over the id in this SQL query. Let's pretend article id 69 is private and contains the number of payouts/bonuses that are to be distributed among employees, so it cannot be viewed on the webpage till Christmas. But you can call,

`https://company.com/notice?id=69;--`

Which would then, in turn, produce the SQL statement:

```SQL
SELECT * FROM USERS WHERE id=1;-- AND role='admin' LIMIT 1;
```

As you can make out from syntax highlighting,
the semicolon in the URL signifies the end of the SQL statement. The logical part of the query is commented out, which will return the article with an id of 69, whether it is set to role admin min or not.

But SQL Injection in the real world wouldn’t be right on the nose like this example, rather it would involve a lot more technicalities. But don’t worry I’ll try my best to explain, what’s going on behind the scenes. But first, let’s look at types of SQL Injections.

```SQL
SELECT * FROM USERS WHERE role='' AND verified='true' OR 1=1 ;
SELECT * FROM EVENTS WHERE id=115 or 1=1 ;

```
