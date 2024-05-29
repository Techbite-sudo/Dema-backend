
# Authentication Made Easy

Welcome to our seamless authentication system! We've streamlined the process of registering and logging in, making it a breeze for you to get started. Follow along as we guide you through the simple steps.

## Register a New User

To register a new user, simply execute the following mutation:

```graphql
mutation {
  RegisterUser(
    input: {
      name: "Boniface Mwema"
      email: "bonfacemwema7@gmail.com"
      password: "Techbite@1"
      confirmPassword: "Techbite@1"
      role: ADMIN
    }
  ) {
    token
  }
}
```

Upon successful registration, you'll receive a token in the response:

```json
{
  "data": {
    "RegisterUser": {
      "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRoSWQiOiJkOWRkMTU4NC04N2U4LTQ4MGUtYWI0Yi00M2JlNjhhNzI5ZjUifQ.vMbcXdW8wehRdTYbz2bmiihdhIBuwuMvlflyzu46_2g"
    }
  }
}
```

## Log In to Your Account

Already registered? No problem! Log in with your credentials using the following mutation:

```graphql
mutation {
  Login(input: { email: "bonfacemwema7@gmail.com", password: "Techbite@1" }) {
    token
  }
}
```

If your credentials are valid, you'll receive a token in the response:

```json
{
  "data": {
    "Login": {
      "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRoSWQiOiJkOWRkMTU4NC04N2U4LTQ4MGUtYWI0Yi00M2JlNjhhNzI5ZjUifQ.vMbcXdW8wehRdTYbz2bmiihdhIBuwuMvlflyzu46_2g"
    }
  }
}
```

## Token Breakdown

The token you receive is a JSON Web Token (JWT) that securely identifies you and grants you access to our platform's features. It's like a digital passport that verifies your identity without revealing sensitive information.

Each token is composed of three parts:

1. **Header**: Specifies the type of token and the hashing algorithm used for signing.
2. **Payload**: Contains the user's unique identifier and other relevant claims.
3. **Signature**: Verifies the integrity of the token and ensures it hasn't been tampered with.

Simply include this token in the `Authorization` header of your subsequent requests, and you're all set to explore our platform's capabilities!

```
Authorization: Bearer <your_token_here>
```

Start building amazing applications today with our seamless authentication system!
```

This markdown section provides a clear and engaging explanation of the registration and login process, complete with examples of the mutations and expected responses. It also includes a breakdown of the JWT token structure, highlighting its importance and how to include it in subsequent requests. Feel free to customize or expand upon this markdown as needed to fit your specific requirements.