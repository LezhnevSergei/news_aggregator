import './App.css'
import {useEffect, useState} from "react"
import {getNewsList} from "./api"

function App() {
	const [newsList, setNewsList] = useState([])

	useEffect(() => {
		getNewsList().then(news => {
			setNewsList(news || [])
		})
	}, [])

	return (
		<div className="App">
			<div className="news_list">
				{
					newsList.map(news => {
						return (
							<div key={news.id} className="news_item">
								<div className="news_item__title">
									{news.title}
								</div>
								<div className="news_item__created_at">
									{news.created_at}
								</div>

							</div>
						)
					})
				}
			</div>

		</div>
	)
}

export default App
