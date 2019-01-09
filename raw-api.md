# Data collection
* users : user,admin,worker
    * POST 
        * /users
    * GET
        * /users
            * get all users
        * /users/{user_id}
            * get users by user_id
    * PUT
        * /users/{user_id}
            * parameter : role(user,admin,worker)
            * update exit user role
    * DELETE
        * /users/{user_id}
* projects
    * GET
        * /projects
            * get all project_name & project_id
            * responce : project_name & project_id
            * role : *
        * /projects/{projectId}
            * get project info by projectId
            * role : *
        * /projects/{projectId}/staus
            * responce : pending,on-going,done
            * role : *
    * POST
        * /projects 
            * create project
            * parameter : project_name,project_id,status_id,data_type,data
            * role : *
    * PUT
        * /projects/{projectId}
            * update exit project
            * role:admin/worker
        * /projects/{projectId}/{status_id}
            * parameter : status
            * role : *
        * /projects/{projectId}/datatype/{data_type}
            * dataType: vedio,image,string
            * role : admin/worker
    * DELETE
         * /projects/{projectId}
             * role : admin

    