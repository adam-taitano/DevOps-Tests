import { Typography, Box, TextField, Paper, Container, Button } from "@mui/material"
import { useState } from 'react';
import {LOGIN} from '../Constants'


interface userInfo {
	name: string,
	password: string
}

const blank_info: userInfo = {
	name: '',
	password: ''
}


const LoginPage = () => {

	let [infoState, setInfoState] = useState(blank_info);
	let [hasSubmit, setHasSubmit] = useState(false);

	const handleSubmit = (event: any) => {
		event.preventDefault();

		console.log(event.target[0].value);
		console.log(event.target[1].value);

		let newInfo: userInfo = {
			name: event.target[0].value,
			password: event.target[1].value
		}

		setInfoState(newInfo);
		setHasSubmit(true);
	}

	
	const SubmitReply = () => {
		if(hasSubmit){
			return(
				<Box mt={10} display='flex' justifyContent='center'>
					<Typography variant='h6'>{LOGIN.SubmitStatement1}{infoState.name}{LOGIN.SubmitStatement2}</Typography>
				</Box>
			)
		}

		const space = ''
		return(<Box>{space}</Box>)
	}
	
	return(
		<>
			<Box mt={30} display='flex' justifyContent='center'>
				<Container maxWidth='xs'>
					<Paper elevation={6}>
						<Box sx={{p:2}} display='flex' justifyContent='center'>
							<Typography variant='h4'>
								{LOGIN.Title}
							</Typography>
						</Box>
						<form onSubmit={handleSubmit}>
							<Box sx={{p:2}} display='flex' justifyContent='center' flexDirection={'column'}>

									<TextField
										required
										id = 'user-name'
										variant = 'filled'
										label = 'User Name'
									/>

								<TextField
									required
									sx={{mt:2}}
									id = 'user-password'
									variant = 'filled'
									label = 'Password'
								/>
								<Button sx={{mt:2}} variant='contained' type='submit' value='Submit'>Submit</Button>
							</Box>
						</form>
					</Paper>
				</Container>
			</Box>

			<SubmitReply/>
		</>
	);
}



export default LoginPage;