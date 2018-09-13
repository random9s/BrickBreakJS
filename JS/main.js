;(() => {
    const canvas = document.getElementById('game-view');
    canvas.setAttribute('width', window.innerWidth);
    canvas.setAttribute('height', window.innerHeight);

    let gameState = {
        score: 0,
        lives: 3,
        ctx: canvas.getContext('2d'),
        render: function () {
        }
    };

    let paddle = {
        isCollider: true,
        ctx: canvas.getContext('2d'),
        pos: {
            x: window.innerWidth/2,
            y: window.innerHeight - 150
        },
        render: function () {
            this.ctx.beginPath ();
            this.ctx.rect(this.pos.x, this.pos.y, 100, 25);
            this.ctx.fillStyle = '#FFFFFF';
            this.ctx.fill ();
            this.ctx.closePath ();
        },
        updatePosition: function () {
            let self = this;
            return function (e) {
                if (e) self.pos.x = e.clientX;
            };
        }
    };
    document.addEventListener('mousemove', paddle.updatePosition());

    let border = {
        ctx: canvas.getContext('2d'),
        isCollider: true,
        offsets: {
            x: 0,
            y: 0
        },
        dimensions: {
            w: window.innerWidth,
            h: window.innerHeight
        },
        render: function () {
            this.ctx.beginPath ();
            this.ctx.rect(this.offsets.x, this.offsets.y, this.dimensions.w, this.dimensions.h);
            this.ctx.fillStyle = '#F4E242';
            this.ctx.fill ();
            this.ctx.closePath ();
        },
        updateDimensions: function () {
            var self = this;
            return function (e) {
                self.dimensions.w = window.innerWidth;
                self.dimensions.h = window.innerHeight;
            }
        }
    };
    window.addEventListener('resize', border.updateDimensions());

    let background = {
        ctx: canvas.getContext('2d'),
        isCollider: false,
        offsets: {
            x: 80,
            y: 45
        },
        dimensions: {
            w: window.innerWidth - 160,
            h: window.innerHeight - 75
        },
        render: function () {
            this.ctx.beginPath ();
            this.ctx.rect(this.offsets.x, this.offsets.y, this.dimensions.w, this.dimensions.h);
            this.ctx.fillStyle = '#000000';
            this.ctx.fill ();
            this.ctx.closePath ();
        },
        updateDimensions: function () {
            var self = this;
            return function (e) {
                self.dimensions.w = window.innerWidth - 160;
                self.dimensions.h = window.innerHeight - 75;
            }
        }

    };
    window.addEventListener('resize', background.updateDimensions());

    const main = () => {
        window.requestAnimationFrame( main );

        background.render();
        paddle.render();
    };

    border.render();
    main ();
})();
