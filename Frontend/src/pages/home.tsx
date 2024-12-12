import { Card, Col, Divider, Button, Row } from 'antd'
import Banner from '../component/banner'
import Game from '../component/game'

function Home() {
  return (
    <>
      <Banner />
      <Divider style={{ border: 1, borderColor: 'gray' }}>Game</Divider>
      <Game />
    </>
  )
}

export default Home