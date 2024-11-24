import React, { useState } from 'react';
import {
    TextField,
    Typography,
    Button,
    Box
} from '@mui/material';
import { makeStyles } from '@mui/styles';
import { postRegister } from '../../api/services/auth.service';
import { useNavigate } from 'react-router-dom';
import { useTranslation } from 'react-i18next';
import { routes } from '../../router/Routes';

const useStyles = makeStyles(() => ({
    paper: {
        padding: 16,
        maxWidth: 400,
        backgroundColor: '#424242', // Slightly lighter grey for contrast
        color: '#ffffff',
        borderRadius: 8,
        boxShadow: '0px 4px 20px rgba(0, 0, 0, 0.5)',
        margin: '0 auto',
    },
    form: {
        display: 'flex',
        flexDirection: 'column',
        gap: 16, // Vertical spacing between fields
    },
    input: {
        '& label': {
            color: '#bdbdbd', // Light grey labels
        },
        '& label.Mui-focused': {
            color: '#ffffff',
        },
        '& .MuiInputBase-root': {
            color: '#ffffff',
        },
        '& .MuiOutlinedInput-root': {
            '& fieldset': {
                borderColor: '#757575', // Grey border
            },
            '&:hover fieldset': {
                borderColor: '#ffffff',
            },
            '&.Mui-focused fieldset': {
                borderColor: '#3f51b5', // Blue when focused
            },
        },
    },
    button: {
        marginTop: 16,
        backgroundColor: '#3f51b5',
        color: '#ffffff',
        '&:hover': {
            backgroundColor: '#303f9f',
        },
    },
}));

interface FormValues {
    firstName: string;
    lastName: string;
    email: string;
    password: string;
    confirmPassword: string;
}

interface Errors {
    email?: string;
    password?: string;
    confirmPassword?: string;
}

export const RegistrationForm: React.FC = () => {
    const { t } = useTranslation();
    const classes = useStyles();
    const navigate = useNavigate();
    const [formValues, setFormValues] = useState<FormValues>({
        firstName: '',
        lastName: '',
        email: '',
        password: '',
        confirmPassword: '',
    });
    const [errors, setErrors] = useState<Errors>({});
    const [isValid, setIsValid] = useState<boolean>(false);

    const validateErrors = (formValues: FormValues) => {
        const errors: Errors = {};
        if (!!formValues.email && (!formValues.email.includes('@') || !formValues.email.includes('.'))) {
            errors.email = t('error.mustBeValidEmail');
        }
        if (!!formValues.password && formValues.password.length < 8) {
            errors.password = t('error.passwordLength');
        }
        if (!!formValues.confirmPassword && formValues.password !== formValues.confirmPassword) {
            errors.confirmPassword = t('error.mismatchPasswords');
        }
        if (!formValues.firstName || !formValues.lastName || !formValues.email || !formValues.password || !formValues.confirmPassword) {
            setIsValid(false);
        }

        if (Object.keys(errors).length === 0) {
            setIsValid(true);
            setErrors({});
        } else {
            setErrors(errors);
            setIsValid(false)
        }
    }

    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        const { name, value } = e.target;
        const updatedFormValues = { ...formValues, [name]: value };
        setFormValues(updatedFormValues);
        validateErrors(updatedFormValues);
    };

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        if (formValues.password !== formValues.confirmPassword) {
            setErrors({ confirmPassword: t('error.mismatchPasswords') });
        } else {
            setErrors({});
            // Submit the form data here


            const formData: FormData = new FormData();
            formData.append("firstName", formValues.firstName);
            formData.append("lastName", formValues.lastName);
            formData.append("email", formValues.email);
            formData.append("password", formValues.password);

            const res = await postRegister(formData);
            if (res?.status === 201) {
                navigate(routes.login);
            }
        }
    };

    const goToLogin = () => {
        navigate(routes.login);
    }

    return (
        <Box className={classes.paper}>
            <Typography variant="h5" align="center" gutterBottom>
                {t('dictionary.register')}
            </Typography>
            <form onSubmit={handleSubmit} className={classes.form}>
                <TextField
                    className={classes.input}
                    label={t('dictionary.firstName')}
                    variant="outlined"
                    fullWidth
                    name="firstName"
                    value={formValues.firstName}
                    onChange={handleChange}
                    required
                />
                <TextField
                    className={classes.input}
                    label={t('dictionary.lastName')}
                    variant="outlined"
                    fullWidth
                    name="lastName"
                    value={formValues.lastName}
                    onChange={handleChange}
                    required
                />
                <TextField
                    className={classes.input}
                    label={t('dictionary.email')}
                    type="email"
                    variant="outlined"
                    fullWidth
                    name="email"
                    value={formValues.email}
                    onChange={handleChange}
                    error={Boolean(errors.email)}
                    helperText={errors.email}
                    required
                />
                <TextField
                    className={classes.input}
                    label={t('dictionary.password')}
                    type="password"
                    variant="outlined"
                    fullWidth
                    name="password"
                    value={formValues.password}
                    onChange={handleChange}
                    error={Boolean(errors.password)}
                    helperText={errors.password}
                    required
                />
                <TextField
                    className={classes.input}
                    label={t('dictionary.confirmPassword')}
                    type="password"
                    variant="outlined"
                    fullWidth
                    name="confirmPassword"
                    value={formValues.confirmPassword}
                    onChange={handleChange}
                    error={Boolean(errors.confirmPassword)}
                    helperText={errors.confirmPassword}
                    required
                />
                <Button
                    type="submit"
                    fullWidth
                    variant="contained"
                    className={classes.button}
                    disabled={!isValid}
                >
                    {t('dictionary.register')}
                </Button>
            </form>

            <Box sx={{ display: 'flex', justifyContent: 'center', marginTop: 2 }}>
                <Typography sx={{ display: 'inline' }}>
                    {t('dictionary.alreadyHaveAccount')}&nbsp;
                </Typography>
                <Typography
                    variant="body1"
                    color="primary"
                    sx={{ cursor: 'pointer', textDecoration: 'underline', display: 'inline' }}
                    onClick={goToLogin}
                >
                    {t('dictionary.here')}
                </Typography>
                <Typography sx={{ display: 'inline' }}>
                    &nbsp;{t('dictionary.toLogin')}
                </Typography>
            </Box>
        </Box>
    );
};
