# Go Web-App

## What makes this project special
### Accomplishments
This is a go Web-App (or website, I have no idea what the difference is), which I started working on for a local homeschool co-op after finding out about the horrors of the system they were using. The project was dropped by the co-op, but I kept working on it for practice.

The thing I am most proud of here is probably the security. I hand-rolled my own Session auth* and ACL middlware. I also added both XSS prevention and form-enforced CSRF mitigation (both via templating and middleware).

*no, the passwords are not in plaintext. They are properly salted and hashed. The session tokens expire properly, are properly cross-site restricted, and are generated using secure and trusted libraries. The least secure thing about this is probably the CSRF storage method (stored in cookie and DOM)

## Running the Website
Just clone the repo; install go, sqlc, echo, and make; and run make in the root directory of this repo

## Next Steps
 - Form Validation

