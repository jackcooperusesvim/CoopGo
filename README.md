# Go Web-App

## What makes this project special
### Accomplishments
This is a go Web-App (or website, I have no idea what the difference is), which I started working on for a local homeschool co-op after finding out about the horrors of the system they were using. The project was dropped by the co-op, but I kept working on it for practice, and maybe to release it as a White-Label SaaS at some point.

The thing I am most proud of here is probably the security. I hand-rolled my own Session auth* and ACL middlware. I also added both XSS prevention and form-enforced CSRF mitigation (both via templating and middleware).

*no, the passwords are not in plaintext. They are properly salted and hashed. The session tokens expire properly, are properly cross-site restricted, and are generated using secure and trusted libraries.

### Anatomy of the Project
The current version of this application is written and organized much better than the original. I kept HTMX on the frontend (I don't like JS) and sqlite3 for the database, but the majority of the backend was completely different. I used Go as the language for the backend, with various supporting modules. These modules were Echo for server routing and middleware, SQLc for SQL type-checking within Go, bcrypt and sha256 for my hand-rolled auth, and templ for HTML templating.

### History of the Project
This is an early re-write for an old attempt I made at building this project. The first version was my first introduction to full-stack web-programming. The stack was HTMX on the frontend, Python/Flask on the Backend, and Sqlite3 for the database. I learned a lot from and since that iteration, though mostly hindsight. For example, I was completely unaware of technologies of HTML templating and Middleware, and had no idea how Cookies worked (causing much frustration when I was working on auth). I didn't learn those from the first version, but the first version taught me about the problems which these technologies solve.

## The Saga: Why this project?

### Why did I start building this.
My Mom used to be the treasurer for the Co-op, so she was in charge of updating the excel sheet which keeps track of all the families, courses, costs, and payments. The Excel sheet was an absolute nightmare, and it would take her hours to change a record. Records could only be edited using macros, meaning that there was no SSOT, and mistakes were very easy to make incredibly hard to fix.

Additionally, the Excel sheet had bad schema. There was one "families" table which would not only keep track of every parent's information, but also the names and grades of all of their children. There were columns titled ChildName1, ChildName2, ChildName3, etc. So every time that the record for largest family was broken at the co-op (It is a Christian Co-op, so this is a regular occurence), new columns would have to be added, and since the only thing that could update the database consistently was macros, all the macros had to be re-written.

I say this in the past tense, but as of writing, the Co-op is still using this system. My efforts on this were initially to replace the Co-ops current system, but those efforts were largely put off. The number and kinds of features the leader of the co-op wanted and expected from this application were of a huge scope; well beyond the capability of an unpaid solo first-time web dev. So now I work on it as a personal project.

