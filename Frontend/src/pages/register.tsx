import { LockOutlined, UserOutlined, PhoneOutlined, MailOutlined } from '@ant-design/icons';
import { Alert, Button, Form, Input, Typography } from 'antd';
import { Link, useNavigate } from 'react-router-dom';
import { useDispatch } from 'react-redux';
import { createUser } from '../store/action/userAction';
import { AppDispatch } from '../store/store';
import { useState } from 'react';


function Register() {
    const dispatch: AppDispatch = useDispatch();
    const navigate = useNavigate();
    const [alert, setAlert] = useState<string>("")

    const onFinish = async (values: any) => {
        try {
            await dispatch(createUser(values)).unwrap();
            setAlert('success')
            setTimeout(() => {
                navigate('/login')
            }, 3000)
        } catch (error: any) {
            setAlert('failed')
        }
    };

    return (
        <>
            {alert === 'success' && (
                <Alert
                    message="Register Success"
                    description="Try your account in Login page"
                    type='success'
                    showIcon
                />
            )}
            {alert === 'failed' && (
                <Alert
                    message="Register Failed"
                    description="Email is already used. Please try again with another email"
                    type='error'
                    showIcon
                />
            )}

            <div style={{ marginTop: 100, marginBottom: 90 }}>
                <Typography.Title level={4} style={{ flex: 1, textAlign: 'center' }}>Register</Typography.Title>
                <div style={{ display: 'flex', justifyContent: 'center', alignItems: 'center' }}>
                    <Form
                        name="register"
                        initialValues={{ remember: true }}
                        style={{ width: 450, padding: 20 }}
                        onFinish={onFinish}
                    >
                        <Form.Item
                            name="name"
                            rules={[{ required: true, message: 'Please input your Name!' }]}
                        >
                            <Input prefix={<UserOutlined />} placeholder="Name" />
                        </Form.Item>
                        <Form.Item
                            name="email"
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
                            name="phoneNumber"
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
        </>
    )
}

export default Register;