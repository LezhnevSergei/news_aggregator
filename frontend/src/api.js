import axios from "axios"

export const getNewsList = () => {
	return axios.get("http://localhost:8080/news").then(r => {
		const newsList = r.data.map(news => {
			return {
				...news,
				created_at: new Date(news.created_at).toUTCString()
			}
		})
		console.log(newsList)
		return newsList
	})
}