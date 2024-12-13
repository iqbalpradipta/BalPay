import { LockOutlined, UserOutlined, PhoneOutlined, MailOutlined } from '@ant-design/icons';
import { Button, Form, Input, Typography } from 'antd';
import { Link } from 'react-router-dom';

function Register() {
    const onFinish = (values: any) => {
        console.log('Received values of form: ', values);
    };

    return (
        <div style={{ marginTop: 100, marginBottom: 90 }}>
            <Typography.Title level={4} style={{ flex: 1, textAlign: 'center' }}>Register</Typography.Title>
            <div style={{ display: 'flex', justifyContent: 'center', alignItems: 'center' }}>
                <Form
                    name="login"
                    initialValues={{ remember: true }}
                    style={{ width: 450, padding: 20 }}
                    onFinish={onFinish}
                >
                    <Form.Item
                        name="Name"
                        rules={[{ required: true, message: 'Please input your Name!' }]}
                    >
                        <Input prefix={<UserOutlined />} placeholder="Name" />
                    </Form.Item>
                    <Form.Item
                        name="Email"
                        rules={[{ required: true, type: 'email', message: 'Please input your Email!' }]}
                    >
                        <Input prefix={<MailOutlined />} placeholder="Email" />
                    </Form.Item>
                    <Form.Item
                        name="password"
                        rules={[{ required: true, message: 'Please input your Password!' }]}
                    >
                        <Input prefix={<LockOutlined />} type="password" placeholder="Password" />
                    </Form.Item>
                    <Form.Item
                        name="PhoneNumber"
                        rules={[{ required: true, message: 'Please input your Phone Number!' }]}
                    >
                        <Input prefix={<PhoneOutlined />} placeholder="Phone Number" />
                    </Form.Item>
                    

                    <Form.Item>
                        <Button block type="primary" htmlType="submit">
                            Register
                        </Button>
                        or <Link to='/login'>Login now!</Link>
                    </Form.Item>
                </Form>
            </div>
        </div>
    )
}

export default Register