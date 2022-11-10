// yes i Do Hate JSon but i have to use it here
const net = require("net");
const { createTransport } = require("nodemailer");

async function sendMail(value) {
    let transporter = createTransport({
        service: "gmail",
        // get data form .env file
        auth: {
            user: "",
            pass: "",
        },
    });

    let mailOptions = {
        from: "praveenswift13@gmail.com",
        to: value.userEmail,
        subject: "click to get verified",
        text: `http://localhost:42069/User/getVerified/${value._id}`,
    };
    try {
        await transporter.sendMail(mailOptions)
        return 'done'
    } catch (error) {
       console.log(error) 
        return 'not-done'
    }
}

const server = net.createServer((c) => {
    c.on("data", async (data) => {
        let result = await sendMail(JSON.parse(data.toString()));
        c.end(result)
    });
});

server.listen(6969, err => err ? console.log(err) : console.log('ws://localhost:6969'))
