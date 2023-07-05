const Application = {
    animationLen: parseInt(getComputedStyle(document.body).getPropertyValue("--length")),

    namespace: undefined,
    namespaces: {
        "calendar": {
            loader: () => 
                new AirDatepicker("#datepicker", {
                    inline: true,
                    dateFormat: "yyyy-MM-dd",
                    buttons: [
                        { content: "Сегодня", onClick: (dp) => dp.selectDate(new Date()) },
                        {
                            content: "Завтра",
                            onClick: (dp) => {
                                const tomorrow = new Date();
                                tomorrow.setDate(tomorrow.getDate() + 1);
                                dp.selectDate(tomorrow);
                            }
                        }
                    ],
                    onSelect: ({ formattedDate }) => Utils.swupRedirect("/calendar/timetable?date=" + formattedDate)
                })
        },
        "calendar-timetable": {
            gallery: {},
            loader: function () {
                const backListener = () => {
                    if (this.gallery.lgOpened) 
                        this.gallery.closeGallery();
            
                    Telegram.WebApp.BackButton.offClick(backListener);
                };
                Telegram.WebApp.BackButton.onClick(backListener);
            
                $(".gallery + button").click((event) => 
                    (this.gallery = lightGallery(event.currentTarget.previousElementSibling, {
                        selector: "li img",
                        controls: false,
                        mobileSettings: {
                            showCloseIcon: true,
                            download: true
                        }
                    })).openGallery()
                );
            },
            button: function () {
                if (this.gallery.lgOpened)
                    this.gallery.closeGallery();

                Utils.swupRedirect(window.location.href + "&edit");
            }
        },
        "calendar-timetable-edit": {
            loader: async () => {
                $("[action=main]").click(function (event) {
                    event.stopPropagation();
            
                    const subject = $(this.closest("li"));
            
                    subject.slideUp(Application.animationLen, () => {
                        if (
                            subject.hasClass("new") ||
                            !subject.find("input[name=name]").val().trim()
                        ) 
                            return subject.remove();

                        subject.find(".option-menu").hide();
            
                        subject.appendTo(subject.parent().prop("id") === "deleted-subjects" ? "#active-subjects" : "#deleted-subjects");
            
                        const option = subject.find(".option-button");
            
                        option.find(".row").toggle();
                        option.toggleClass("inactive");
            
                        subject.slideDown(Application.animationLen);
                    });
                });
            
                const template = $("#template");
                document.getElementById("add").addEventListener("click", () =>
                    template
                        .clone(true, true)
                        .removeProp("id")
                        .appendTo("#active-subjects")
                        .slideDown(Application.animationLen)
                );

                const checkHomework = subject => {
                    const name = subject.querySelector("input[name=name]").value.trim().toLowerCase();
            
                    const firstTextarea = subject.querySelector("textarea");
            
                    for (const li of document.querySelectorAll("#active-subjects > li")) {
                        if (li === subject) 
                            continue;

                        if (li.querySelector("input[name=name]").value.trim().toLowerCase() === name) {
                            const secondTextarea = li.querySelector("textarea");
                            if (secondTextarea.value || li.querySelector(".gallery").childElementCount)
                                $(firstTextarea.closest("li")).hide();
                            else 
                                $(secondTextarea.closest("li"))[
                                    firstTextarea.value || subject.querySelector(".gallery").childElementCount ? "hide" : "show"
                                ]();

                            return;
                        }
                    }

                    $(firstTextarea.closest("li")).show();
                };
            
                $("input[name=name]").change(function () {
                    const button = this.closest(".option-button");
                    const subject = $(button.parentElement);
            
                    if (!subject.hasClass("new")) {
                        const clone = subject.clone(true, true);
                        clone.find("input[name=name]").val(this.defaultValue);
                        clone.find("[action=main]").click();
            
                        subject.addClass("new");
                    }
            
                    const name = this.value.trim().toLowerCase();
                    const menu = button.nextElementSibling;
            
                    for (const deleted of document.querySelectorAll("#deleted-subjects > li")) 
                        if (deleted.querySelector("input[name=name]").value.toLowerCase() === name) {
                            button.querySelector("[action=main]").click();
                            deleted.querySelector("[action=main]").click();

                            return $(menu).slideUp(Application.animationLen);
                        }
            
                    const original = [...document.getElementById("default-subjects").children].find(el => el.value.toLowerCase() === name);
                    if (original) {
                        button.querySelector("input[name=name]").value = original.value;
        
                        menu.querySelector("input[name=classroom]").value = original.getAttribute("classroom");
                        menu.querySelector("input[name=teacher]").value = original.getAttribute("teacher");
                    }

                    checkHomework(button.parentElement);
                });
            
                $("textarea").change(function () {
                    checkHomework(this.closest("#active-subjects > li"));
                });
            
                $("[action=add]").click(function () {
                    this.nextElementSibling.click();
                });
            
                $("[action=remove]").click(function () {
                    const subject = this.closest("#active-subjects > li");
            
                    this.parentElement.remove();
            
                    checkHomework(subject);
                });
            
                const templateGallery = $("#template-gallery");
                $("input[type=file]")
                    .change(async function () {
                        if (this.files && this.files.length) {
                            const gallery = $(this).siblings(".gallery");

                            const promises = [];
                            for (const file of this.files) {
                                const reader = new FileReader();
                    
                                promises.push(new Promise(resolve => {
                                    reader.onload = event => resolve(event.target.result);
                                    reader.readAsDataURL(file);
                                }));
                            }
                
                            for (const file of await Promise.all(promises)) {
                                const clone = templateGallery.clone(true, true).removeAttr("id");
                                clone.find("> img").prop("src", file);
                                clone.appendTo(gallery).show();
                            }
                
                            checkHomework(this.closest("#active-subjects > li"));
                        }
                    })
                    .click(function () {
                        this.value = null;
                    });
            
                new Sortable(document.getElementById("active-subjects"), {
                    animation: Application.animationLen,
                    delay: Application.animationLen / 2,
                    filter: ".option-menu",
                    preventOnFilter: false
                });
            },
            button: async () => {
                const form = document.querySelector("form");
                if (!form.checkValidity()) 
                    return;
            
                Telegram.WebApp.MainButton.showProgress();
            
                const timetable = {
                    "date": Utils.getParam("date"),
                    "offset": form.querySelector("input[name=offset]")
                        .value
                        .split(":")
                        .reduce((total, val, i) => total + (i ? +val : 60 * +val), 0),
                    "subjects": [],
                    "note": form.querySelector("input[name=note]").value.trim()
                };
            
                for (const subject of form.querySelectorAll("#active-subjects > li")) {
                    const result = {
                        ...Utils.formData(subject),
                        "homework": {
                            "text": subject.querySelector("textarea[name=homework]").value.trim(),
                            "images": [...subject.querySelectorAll(".gallery > li > img")].map(img => img.src)
                        }
                    };

                    if (!result.homework.text && !result.homework.images.length)
                        delete result.homework;

                    timetable.subjects.push(result);
                }

                const notification = form.querySelector("input[name=notification]").checked ? "&notification" : "";
            
                await Utils.ajax("/api/timetable?type=date" + notification, "put", timetable);
            
                Telegram.WebApp.MainButton.hideProgress();
            
                window.history.back();
            }
        },
    
        "registration": {
            button: async () => {
                const form = document.querySelector("form");
                if (!form.checkValidity()) 
                    return;
            
                Telegram.WebApp.MainButton.showProgress();
            
                await Utils.ajax("/api/class", "post", Utils.formData(form));
            
                Telegram.WebApp.MainButton.hideProgress();

                window.onbeforeunload = null;
                location.reload();
            },
        },
    
        "class": {
            button: () => Utils.swupRedirect(window.location.href + "?edit")
        },
        "class-edit": {
            loader: () => {
                $("[action]").click(function () {
                    if (this.classList.contains("inactive"))
                        return;

                    const $this = $(this);

                    switch (this.getAttribute("action")) {
                        case "promote": {
                            const input = $this.siblings("input");
                            const prevRole = +input.val();
                            const role = input.val(prevRole + 1).val();
                            if (role === input.prop("max"))
                                $this.addClass("inactive");

                            $this.siblings("[role=" + prevRole + "]").hide();
                            $this.siblings("[role=" + role + "]").show();
                            
                            return $this.siblings("[action=demote]").removeClass("inactive");
                        }
                        case "demote": {
                            const input = $this.siblings("input");
                            const prevRole = input.val();
                            const role = input.val(prevRole - 1).val();
                            if (role === input.prop("min"))
                                $this.addClass("inactive");

                            $this.siblings("[role=" + prevRole + "]").hide();
                            $this.siblings("[role=" + role + "]").show();
                            
                            return $this.siblings("[action=promote]").removeClass("inactive");
                        }
                        case "accept": 
                            return $this.hide().closest("li").appendTo("#students .option-menu");
                        case "delete": {
                            const li = $this.closest("li");
                            if (li.closest("#applications").length)
                                return li.remove();
                            
                            $this.siblings("[action=accept]").show();
                            li.appendTo("#applications .option-menu");
                        }
                    }
                });
            },
            button: async () => {
                const form = document.querySelector("form");
                if (!form.checkValidity()) 
                    return;
            
                Telegram.WebApp.MainButton.showProgress();

                const students = [];
                for (const student of form.querySelectorAll(".option-menu > li")) 
                    students.push({ ...Utils.formData(student), active: !!student.closest("#students") });

                await Utils.ajax("/api/class", "put", students);
            
                Telegram.WebApp.MainButton.hideProgress();
            
                window.history.back();
            }
        },
    
        "control-timetable": {
            loader: async () => {
                $("[action]").click(function (event) {
                    event.stopPropagation();

                    $(this.closest("li")).slideUp(Application.animationLen, function () {
                        this.remove();
                    });
                });
            
                const template = $("#template");
                document.getElementById("add").addEventListener("click", function () {
                    template
                        .clone(true, true)
                        .removeProp("id")
                        .appendTo("#subjects")
                        .slideDown(Application.animationLen);
                });
            
                $("input[name=name]").change(function() {
                    const button = this.closest(".option-button");
            
                    const name = this.value.trim().toLowerCase();
                    const menu = button.nextElementSibling;
            
                    const original = [...document.getElementById("default-subjects").children].find(el => el.value.toLowerCase() === name);
                    if (original) {
                        button.querySelector("input[name=name]").value = original.value;
        
                        menu.querySelector("input[name=classroom]").value = original.getAttribute("classroom");
                        menu.querySelector("input[name=teacher]").value = original.getAttribute("teacher");
                    }
                });
            
                new Sortable(document.getElementById("subjects"), {
                    animation: Application.animationLen,
                    delay: Application.animationLen / 2,
                    filter: ".option-menu",
                    preventOnFilter: false
                });
            },
            button: async () => {
                const form = document.querySelector("form");
                if (!form.checkValidity()) 
                    return;
            
                Telegram.WebApp.MainButton.showProgress();
            
                const timetable = {
                    "day": +Utils.getParam("day"),
                    "offset": form.querySelector("input[name=offset]")
                        .value
                        .split(":")
                        .reduce((total, val, i) => total + (i ? +val : 60 * +val), 0),
                    "subject_length": +form.querySelector("input[name=subject_length]").value,
                    "subject_break": +form.querySelector("input[name=subject_break]").value
                }

                timetable.subjects = [...form.querySelectorAll("#subjects > li")].map(Utils.formData);

                await Utils.ajax("/api/timetable?type=day", "put", timetable);
            
                Telegram.WebApp.MainButton.hideProgress();
            
                window.history.back();
            }
        },
        "control-subjects": {
            loader: () => {
                $("[action]").click(function () {
                    $(this.closest("li")).slideUp(Application.animationLen, function () {
                        this.remove();
                    });
                });
            
                const template = $("#template");
                document.getElementById("add").addEventListener("click", () =>
                    template
                        .clone(true, true)
                        .removeProp("id")
                        .appendTo("#subjects")
                        .slideDown(Application.animationLen)
                );
            
                $("input[name=name]").change(function () {
                    const val = this.value.trim().toLowerCase();
            
                    for (const el of document.querySelectorAll("#subjects > li input[name=name]")) {
                        if (el === this) 
                            continue;
            
                        if (el.value.trim().toLowerCase() === val) 
                            return $(this.closest("li")).slideUp(Application.animationLen, function () {
                                this.remove();
                            });
                    }
                });
            },
            button: async () => {
                const form = document.querySelector("form");
                if (!form.checkValidity()) 
                    return;
            
                Telegram.WebApp.MainButton.showProgress();
            
                await Utils.ajax("/api/subjects", "put", [...form.querySelectorAll("#subjects > li")].map(Utils.formData));
            
                Telegram.WebApp.MainButton.hideProgress();
            
                window.history.back();
            }
        },
        "control-holidays": {
            loader: () => {
                const template = $("#template");

                const dp = new AirDatepicker("#datepicker", {
                    isMobile: true,
                    dateFormat: "yyyy-MM-dd",
                    range: true,
                    multipleDatesSeparator: " - ",
                    buttons: [
                        {
                            content: "Выбрать",
                            onClick: () => {
                                const input = dp._getInputValue(dp.locale.dateFormat);
                                if (!input) 
                                    return;
            
                                const holiday = template.clone(true, true).removeProp("id");
                                holiday.appendTo(".menu ul").slideDown(Application.animationLen).find("input").val(input);
            
                                dp.hide();
                                dp.clear();
                            }
                        }
                    ]
                });
            
                document.getElementById("add").addEventListener("click", () => dp.show());

                const backListener = () => {
                    dp.hide();

                    Telegram.WebApp.BackButton.offClick(backListener);
                };
                Telegram.WebApp.BackButton.onClick(backListener);
            
                $("[action]").click(function () {
                    $(this.closest("li")).slideUp(Application.animationLen, function() {
                        this.remove();
                    });
                });
            },
            button: async () => {
                Telegram.WebApp.MainButton.showProgress();

                const holidays = [];
                for (const input of document.querySelectorAll("form input[name=holiday]")) {
                    const days = input.value.split(" - ");

                    holidays.push({
                        start_date: days[0],
                        end_date: days[1]
                    });
                }

                await Utils.ajax("/api/holidays", "put", holidays);
            
                Telegram.WebApp.MainButton.hideProgress();
            
                window.history.back();
            }
        }
    },

    swup: undefined,

    init() {
        this.initHeader();
        this.initSwup();
        this.initTelegram();

        window.onbeforeunload = () => Telegram.WebApp.close();

        this.loadNamespace();
    },
    initHeader() {
        $("[destination]").click(function () {
            const $this = $(this);
            if ($this.hasClass("selected")) 
                return;
        
            $this.siblings(".selected").removeClass("selected");
            $this.addClass("selected");
        
            Utils.swupRedirect($this.attr("destination"));
        });
    },
    initSwup() {
        this.swup = new Swup({
            animateHistoryBrowsing: true,
            cache: false
        });
        
        this.swup.on("contentReplaced", () => this.loadNamespace());
        this.swup.on("transitionStart", () => {
            Telegram.WebApp.MainButton.hide();
            Telegram.WebApp.BackButton.hide();
        });
    },
    initTelegram() {
        Telegram.WebApp.BackButton.onClick(() => window.history.back());
        Telegram.WebApp.MainButton.onClick(() => this.namespace.button());

        Telegram.WebApp.expand();
    },

    loadNamespace() {
        $(".option-button").click(function (event) {
            const target = event.target;
    
            if (target instanceof HTMLInputElement || target.closest(".inactive")) 
                return;
    
            const menu = $(this).siblings(".option-menu");
            if (menu.children().length) 
                menu.stop().slideToggle(Application.animationLen);
        });
    
        $("[swup-page]").click(function () {
            Utils.swupRedirect(this.getAttribute("swup-page"));
        });
    
        const swup = document.getElementById("swup");
    
        Telegram.WebApp.BackButton[swup.hasAttribute("no-return") ? "hide" : "show"]();
    
        const text = swup.getAttribute("main-button");
        if (text) {
            Telegram.WebApp.MainButton.text = text;
            Telegram.WebApp.MainButton.show();
        } else 
            Telegram.WebApp.MainButton.hide();

        this.namespace = this.namespaces[swup.getAttribute("namespace")];
        if (this.namespace && this.namespace.loader) 
            this.namespace.loader();
    }
};

const Utils = {
    swupRedirect(url) {
        Application.swup.loadPage({ url });
    },
    
    async ajax(url, method, content) {
        if (!content) 
            return await fetch(url, { method });
    
        return await fetch(url, {
            method,
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify(content)
        });
    },

    formData(el) {
        return [...el.querySelectorAll("input")].reduce(
            (acc, input) => ({ ...acc, [input.name]: input.type === "number" ? +input.value : input.value.trim() }),
        {});
    },
    
    getParam(param) {
        return new URL(window.location.href).searchParams.get(param);
    }
};