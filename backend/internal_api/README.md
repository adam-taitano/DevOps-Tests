# Build Backend:
> Requires .env file in base directory
> 
> [Example .env file](../example.env) 
>
> In backend directory: 
> ```
> cargo run
> ```


# API Routes:

## GET api/
> ### Request Example:
> > No Request Parameters
> ### Response Example:
> > ```
> > {
> >     "active": true
> > }
> > ```


## GET api/users
> ### Request Example:
> > No Request Parameters
> ### Response Example:
> > ```
> > [
> >     {
> >         "id": 1,
> >         "username": "testuser1"
> >     },
> >     {
> >         "id": 2,
> >         "username": "testuser2"
> >     }
> > ]
> > ```