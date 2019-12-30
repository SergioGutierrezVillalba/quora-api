var express = require('express');
var quora = require('quora-api');
var User = require('./model/User');


var app = express();

app.get('/user', async function(req, res){

    let userLink = 'profile/' + req.headers.user;

    // let userLink = 'profile/Sergio-Gutiérrez-5';
    let userInfo;

    await quora(userLink)
        .then(user => {

            let user_name = user.$('.user')['0'].children[0].data || null;
            let bio = user.$('div.ProfileDescription').text() || null;
            let answers = user.$('.list_count').eq(0).text() || null;
            let questions = user.$('.list_count').eq(1).text() || null;
            let shares = user.$('.list_count').eq(2).text() || null;
            let publications = user.$('.list_count').eq(3).text() || null;
            let followers = user.$('.list_count').eq(4).text() || null;
            let following = user.$('.list_count').eq(5).text() || null;
            // let monthlyViews = user.$('.total_count').eq(0).text() || null;
			// let totalViews = user.$('.total_count').eq(1).text() || null;

            let work = user.$('.WorkCredentialListItem').find('span.UserCredential').text();
            let studies = user.$('.SchoolCredentialListItem').find('span.UserCredential').text();
            let location = user.$('.LocationCredentialListItem').find('span.UserCredential').text();
            
            let contentViews = user.$('.ContentViewsAboutListItem');
            let monthlyViews = contentViews.find('span.main_text').text();
            let totalViews = contentViews.find('span.detail_text').text();

            let userInfo = new User(
                user_name, 
                bio, 
                answers, 
                questions, 
                shares, 
                publications, 
                followers, 
                following, 
                monthlyViews, 
                totalViews,
                location,
                work,
                studies
            )

            return res.status(200).send(userInfo)

        })
        .catch((err) => {
            console.log(err)
            res.status(500).send({userInfo: 'error'})
        })
    
})


app.listen(3000)