import { Carousel, Image} from 'antd'
import banner from '../mocks/banner.json'

function Banner() {
    return (
        <>
            <Carousel autoplay autoplaySpeed={3000} centerMode centerPadding='250px' infinite slidesToShow={1}>
                {banner.map((value, index) => (
                    <div key={index}>
                        <Image
                            src={value.LinkImage}
                            style={{
                                width: 800,
                                height: 350,
                            }}
                        />
                    </div>
                ))}
            </Carousel>
        </>
    )
}

export default Banner