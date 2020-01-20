# Pgcrypto Demos
These demos utilize the `pgcrytpo` extension which is provided by `postgresql` version `11`. In these three demos, there are examples of :
  - Hashing and salting passwords
  - Storing usernames and **hashed** passwords
  - Performing queries for user authentication

### Background ###
Hashing passwords is extremely important whenever we deal with user data. If a malicious attacker was somehow able to gain access to usernames and passwords and the passwords were in plaintext, the accounts would be completely compromised. Hashes are **one-way** functions. So, if an attacker gets hold of a hash, they cannot "decrypt" it to find the password. Instead, the only thing they can do is hash other "likely" passwords and see if the hashes match. This makes life harder for the attacker and a total compromise less likely.

When a user provides their credentials (username and password), we store their information in a table called `users`. We **never** store the plaintext password in the database--instead, we store the **hash**. The next time a user logs in, we hash their password, see if it matches, and then grant them authorization.

### Dependencies ###
These demos require the interactive `postgresql` environment, `psql`. Follow the instructions for your respective system:
  - https://postgresql.org/download
 
# Running the Demos
1. Make sure you have the `pgcrypto_demos` repository downloaded on your machine.
2. Move into your interactice `psql` environment by doing: `sudo -u postgres -i`
3. Do a recurisve copy to move the `pgcrypto_demos` from your local machien into your local `psql` environment. For example:
        ```
        $ cp -r /home/rahultoppur/pgcrypto_demos/ .
        ```
        This moves the `pgcrypto_demos` directory into your local directory, `.`
4. From here, you can manually open up each file walk through them for yourself. Read the comments to have a better idea of what's going on, or run each demo individually to observe the output. To open each demo in an editor, `cd` into the `pgcrypto_demos` directory and use `vim` or an editor of your choosing. Use `Control-D` to exit the `psql` environment.

    Run each demo by doing: `psql -f <filename>.psql`, e.g. `psql -f hash_salt_demo.psql`. 
5. The demos are broken up into three parts:
  -`hash_salt_demo` : How to salt and hash usernames and passwords
  -`store_email_passwd_demo` : How to store emails and **hashed** passwords in a `USERS` table
  -`user_login_lookup_demo` : Performing queries to match hashses and authenticate users.

# Helpful Resources
  - https://www.postgresql.org/docs/8.3/pgcrypto.html
  - https://chartio.com/resources/tutorials/how-to-list-databases-and-tables-in-postgresql-using-psql/
  - https://www.postgresql.org/docs/

