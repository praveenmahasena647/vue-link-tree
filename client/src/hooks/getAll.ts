import axios from 'axios'

async function getAllUsers() {
    try {
        let {data}= await axios.get('http://localhost:42069/User/')
        console.log(data)
        return data
    } catch (error) {
        console.log(error)
        return 10
    }
}

export { getAllUsers }
