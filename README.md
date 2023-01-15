# task-5-vix-btpns-Achmad-kautsar-masuara

## Endpoint Users
## ```POST : /users/register```
* header
    ```
        None
    ```
* Body  ( JSON )     
    ```
        {
            "username"  :string,
            "email"     :email,
            "password"  :string
        }
    ```

## ```POST : /users/login```
* header
    ```
        None
    ```
* Body  ( JSON )     
    ```
        {
            "email"     :email,
            "password"  :string
        }
    ```

## ```PUT : /users/:userId```
* header
    ```
        None
    ```
* Body  ( JSON )     
    ```
        {
            "username"  :string,
            "email"     :email,
            "password"  :string
        }
    ```

## ```DELETE : /users/:userId```
* header
    ```
        None
    ```
* Body 
    ```
        None
    ```

## Endpoint Photos
## ```POST : /Photos```
* header
    ```
        access_token
    ```
* Body  ( JSON )     
    ```
        {
            "title"     :string,
            "caption"   :string,
            "photo_url" :string
        }
    ```

## ```GET : /Photos```
* header
    ```
        access_token
    ```
* Body     
    ```
        None
    ```

## ```PUT : /:PhotoId```
* header
    ```
        access_token
    ```
* Body  ( JSON )     
    ```
        {
            "title"     :string,
            "caption"   :string,
            "photo_url" :string
        }
    ```

## ```DELETE : /:PhotoId```
* header
    ```
        access_token
    ```
* Body      
    ```
        None
    ```