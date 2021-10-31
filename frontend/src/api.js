import axios from "axios"

export const getNewsList = () => {
	return axios.get("http://localhost:8080/news").then(r => {
		return  r.data?.map(news => {
			return {
				...news,
				created_at: new Date(news.created_at).toUTCString()
			}
		})
	})
}

export const getNewsByTitle = (search) => {
	return axios.get(`http://localhost:8080/news?search=${search}`).then(r => {
		return r.data?.map(news => {
			return {
				...news,
				created_at: new Date(news.created_at).toUTCString()
			}
		})
	})
}