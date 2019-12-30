class User {

    constructor(
        username = '',
        bio = '', 
        answers = '', 
        questions = '', 
        shares = '', 
        publications = '', 
        followers = '', 
        following = '', 
        monthlyViews = '', 
        totalViews = '',
        location = '',
        work = '',
        studies = ''
    ){
        this.username = username;
        this.bio = bio;
        this.answers = answers;
        this.questions = questions;
        this.shares = shares;
        this.publications = publications;
        this.followers = followers;
        this.following = following;
        this.monthlyViews = monthlyViews;
        this.totalViews = totalViews;
        this.credentials = {};
        this.credentials['location'] = location;
        this.credentials['work'] = work;
        this.credentials['studies'] = studies;
    }

}

module.exports = User;