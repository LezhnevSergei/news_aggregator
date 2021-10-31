import './App.css'
import {useEffect, useState} from "react"
import {getNewsByTitle, getNewsList} from "./api"

function App() {
	const [newsList, setNewsList] = useState([])
	const [search, setSearch] = useState("")

	useEffect(() => {
		getNewsList().then(news => {
			setNewsList(news || [])
		})
	}, [])

	const searchTextHandler = (e) => {
		setSearch(e.target.value)
	}

	const searchHandler = () => {
		getNewsByTitle(search).then(news => {
			setNewsList(news || [])
			setSearch("")
		})
	}

	return (
		<div className="App">
			<div className="news_search">
				<input
					type="text"
					className="news_search__input"
					value={search}
					onChange={e => searchTextHandler(e)}
				/>
				<button
					className="news_search__button"
					onClick={searchHandler}
				>
					Поиск
				</button>
			</div>
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
