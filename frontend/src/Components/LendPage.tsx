import { Typography, Box } from "@mui/material";
import { LEND } from "../Constants";

const LendPage = () => {
	return(
		<>
			<Box mt={15} display='flex' justifyContent='center'>
				<Typography variant='h3'>
					{LEND.welcome}
				</Typography>
			</Box>

			<Box display='flex' justifyContent='center'>
				<Typography variant='h5'>
					{LEND.subtitle}
				</Typography>
			</Box>
		</>
	);
}



export default LendPage;