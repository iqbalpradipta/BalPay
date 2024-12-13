import { LockOutlined, UserOutlined } from '@ant-design/icons';
import { Button, Checkbox, Form, Input, Flex, Typography } from 'antd';
import { Link } from 'react-router-dom';

function Login() {
    const onFinish = (values: any) => {
        console.log('Received values of form: ', values);
    };

    return (
        <>
            <div style={{marginTop: 100, marginBottom: 150}}>
                <Typography.Title level={4} style={{ flex: 1, textAlign: 'center' }}>Login</Typography.Title>
                <div style={{ display: 'flex', justifyContent: 'center', alignItems: 'center' }}>
                    <Form
                        name="login"
                        initialValues={{ remember: true }}
                        style={{ width: 450, padding: 20 }}
                        onFinish={onFinish}
                    >
                        <Form.Item
                            name="Email"
                            rules={[{ required: true, type: 'email', message: 'Please input your Email!' }]}
                        >
                            <Input prefix={<UserOutlined />} placeholder="Email" />
                        </Form.Item>
                        <Form.Item
                            name="password"
                            rules={[{ required: true, message: 'Please input your Password!' }]}
                        >
                            <Input prefix={<LockOutlined />} type="password" placeholder="Password" />
                        </Form.Item>
                        <Form.Item>
                            <Flex justify="space-between" align="center">
                                <Form.Item name="remember" valuePropName="checked" noStyle>
                                    <Checkbox>Remember me</Checkbox>
                                </Form.Item>
                                <a href="">Forgot password</a>
                            </Flex>
                        </Form.Item>

                        <Form.Item>
                            <Button block type="primary" htmlType="submit">
                                Log in
                            </Button>
                            or <Link to='/register'>Register now!</Link>
                        </Form.Item>
                    </Form>
                </div>
            </div>
        </>
    )
}

export default Login