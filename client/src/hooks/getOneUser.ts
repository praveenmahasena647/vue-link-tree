import axios from 'axios'
async function getHim(person: string | any) {
    try {
       let {data}=await axios.get(`http://localhost:42069/User/${person}`) 
        console.log(data)
    } catch (error) {
        console.log('there was an err')
       console.log(error) 
    }
}

export { getHim }
