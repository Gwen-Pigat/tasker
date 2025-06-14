import { user,error,API_URL } from '$lib'
import { get } from 'svelte/store'


console.log(get(user))

export async function resetUser(){
    user.set({})
}

export async function fetchAPI(
    path: string, 
    method: string,
    data?:any
):Promise<any>{
    let result  
    const headers = new Headers()   
    const currentUser:any = get(user)
    if(currentUser.token !== undefined){
        headers.append("Authorization", "Bearer "+currentUser.token)
    }
    const options:any = {
        method: method,
        redirect: "follow",
        headers: headers
    }
    if(method === "POST" && data){
        options.body = data
    }
    try{
        const response = await fetch(API_URL+path, options)
        console.log(API_URL+path, options)
        result = await response.json()
        if(!response.ok){
            let error = "Error request"
            if(result.error){
                error = result.error
            }
            if(response.status === 401){
                return {
                    "error": error
                }
            }
            throw error
        }
        
    } catch(err:any){
        return {
            "error": err
        }
    }
    return result
}