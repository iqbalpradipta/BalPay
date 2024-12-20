import { Card, Col, Button, Row } from 'antd'
import game from '../mocks/game.json'
import { Link, useParams } from 'react-router-dom'
import { useDispatch, useSelector } from 'react-redux'
import { useEffect } from 'react'
import { gameFetch } from '../store/action/gameAction'
import { AppDispatch, RootState } from '../store/store'

const { Meta } = Card
function Game() {
    const dispatch = useDispatch() as AppDispatch;
    const games = useSelector((state: RootState) => state.games);
    const { id } = useParams()
    useEffect(() => {
        dispatch(gameFetch());
    }, [dispatch, gameFetch]);

    return (
        <>
            {games.data && games.data.length > 0 && (
                <Row gutter={{ xs: 8, sm: 16, md: 24, lg: 32 }}>
                    {games?.data?.map((value, index) => (
                        <Col key={index} xs={24} sm={12} md={8} lg={6}>
                            <div style={{ padding: 20 }}>
                                <Card
                                    hoverable
                                    style={{ width: 250 }}
                                    cover={<img alt={value.name as string} src={value.icon as string} />}
                                >
                                    <Meta title={value.name} description={value.description} />
                                    <Button variant="outlined" color="default" style={{ marginTop: 20 }}>
                                        <Link to="/product/productType">Beli Sekarang!</Link>
                                    </Button>
                                </Card>
                            </div>
                        </Col>
                    ))}
                </Row>
            )}
        </>
    );
}

export default Game