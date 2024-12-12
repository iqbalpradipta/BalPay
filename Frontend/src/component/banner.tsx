import { Carousel, Image} from 'antd'
import banner from '../mocks/banner.json'

function Banner() {
    return (
        <>
            <Carousel autoplay autoplaySpeed={5000} centerMode centerPadding='400px' infinite slidesToShow={1}>
                {banner.map((value, index) => (
                    <div key={index}>
                        <Image
                            src={value.LinkImage}
                            style={{
                                width: 500,
                                height: 220,
                            }}
                        />
                    </div>
                ))}
            </Carousel>
        </>
    )
}

export default Banner