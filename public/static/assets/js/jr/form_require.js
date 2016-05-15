﻿define(['jr/core'],function(j6) {
    jr.extend({
        form: {
            getData: function (a) {
                var b = '';
                var c = document.forms[a || 0];
                return jr.json.toQueryString(c)
            }, asyncSubmit: function (a, b) {
                var c = document.forms[a || 0];
                var d = document.getElementById('$async_ifr');
                if (!d) {
                    try {
                        d = document.createElement('<iframe name="$async_ifr">')
                    } catch (ex) {
                        d = document.createElement('iframe');
                        d.setAttribute('name', '$async_ifr')
                    }
                    d.setAttribute('id', '$async_ifr');
                    if (!b) {
                        d.style.cssText = 'display:none'
                    } else {
                        d.style.cssText = 'width:600px;height:400px'
                    }
                    document.body.insertBefore(d, document.body.firstChild)
                }
                if (c.getAttribute('target') != d.name) {
                    c.setAttribute('target', d.getAttribute('name'))
                }
                c.submit()
            }
        }
    });
    jr.extend({
        validator: {
            setTip: function (e, a, b, c) {
                if (b) {
                    var d = e.getAttribute('summary');
                    if (d) {
                        d = jr.toJson(d);
                        if (d[b]) {
                            c = d[b]
                        }
                    }
                }
                if (e.getAttribute('tipin')) {
                    var f = document.getElementById(e.getAttribute('tipin'));
                    if (f) {
                        if (f.className.indexOf('validator') == -1) {
                            f.className += ' validator'
                        }
                        f.innerHTML = '<span class="' + (a ? 'valid-right' : 'valid-error') + '"><span class="msg">' + c + '</span></span>';
                        return false
                    }
                }
                var g = e.getAttribute('validate_id');
                var h = document.getElementById(g);
                if (!h) {
                    h = document.createElement('DIV');
                    h.id = g;
                    h.className = 'validator';
                    var i = jr.getPosition(e);
                    h.style.cssText = 'position:absolute;left:' + (i.right + document.documentElement.scrollLeft) + 'px;top:' + (i.top + document.documentElement.scrollTop) + 'px';
                    document.body.appendChild(h)
                }
                h.innerHTML = '<span class="' + (a ? 'valid-right' : 'valid-error') + '"><span class="msg">' + c + '</span></span>'
            }, removeTip: function (e) {
                if (e.getAttribute('tipin')) {
                    var a = document.getElementById(e.getAttribute('tipin'));
                    if (a) {
                        a.innerHTML = '';
                        return false
                    }
                }
                var b = document.getElementById(e.getAttribute('validate_id'));
                if (b) {
                    document.body.removeChild(b)
                }
            }, result: function (a) {
                if (a) {
                    var b = true;
                    var c = document.getElementById(a);
                    jr.each(jr.dom.getsByClass(c, 'ui-validate'), function (i, e) {
                        if (b) {
                            if (e.getAttribute('tipin')) {
                                if (jr.$(e.getAttribute('tipin')).innerHTML.indexOf('valid-error') != -1) {
                                    b = false
                                }
                            } else {
                                e = document.getElementById(e.getAttribute('validate_id'));
                                if (jr.dom.getsByClass(e, 'valid-error').length != 0) {
                                    b = false
                                }
                            }
                        }
                    });
                    return b
                } else {
                    return jr.dom.getsByClass(document, 'valid-error').length == 0
                }
            }, init: function () {
                var f = j6;
                if (!f) {
                    alert('未引用核心库!');
                    return false
                }
                var g = document.getElementsByClassName('ui-validate');
                var h;
                for (var i = 0; i < g.length; i++) {
                    h = g[i].getAttribute('validate_id');
                    while (h == null) {
                        h = g[i].id;
                        if (h && h != '') {
                            h = 'validate_item_' + h
                        } else {
                            h = 'validate_item_' + parseInt(Math.random() * 1000).toString()
                        }
                        if (document.getElementById(h) != null) {
                            h = null
                        } else {
                            g[i].setAttribute('validate_id', h)
                        }
                    }
                    var j = new Array();
                    if (g[i].onblur) {
                        j[j.length] = g[i].onblur
                    }
                    if (g[i].getAttribute('isnumber') == 'true') {
                        g[i].style.cssText += 'ime-mode:disabled';
                        var k = (function (a, e) {
                            return function () {
                                if (/\D/.test(e.value)) {
                                    e.value = e.value.replace(/\D/g, '')
                                }
                                e.value = e.value.replace(/^0([0-9])/, '$1')
                            }
                        })(this, g[i]);
                        jr.event.add(g[i], 'keyup', k);
                        jr.event.add(g[i], 'change', k)
                    }
                    if (g[i].getAttribute('regex')) {
                        var k = (function (b, e) {
                            return function () {
                                var a = new RegExp();
                                a.compile(e.getAttribute('regex'));
                                if (!a.test(e.value)) {
                                    b.setTip(e, false, 'regex', '输入不正确')
                                } else {
                                    b.removeTip(e)
                                }
                            }
                        })(this, g[i]);
                        j[j.length] = k
                    } else {
                        if (g[i].getAttribute('isrequired') == 'true' || g[i].getAttribute('required') == 'true') {
                            var k = (function (a, e) {
                                return function () {
                                    if (e.value.replace(/\s/g, '') == '') {
                                        a.setTip(e, false, 'required', '该项不能为空')
                                    } else {
                                        a.removeTip(e)
                                    }
                                }
                            })(this, g[i]);
                            j[j.length] = k
                        }
                        if (g[i].getAttribute('length')) {
                            var k = (function (d, e) {
                                return function () {
                                    var a = e.getAttribute('length');
                                    var b = /\[(\d*),(\d*)\]/ig;
                                    var c = parseInt(a.replace(b, '$1')), l_e = parseInt(a.replace(b, '$2'));
                                    if (e.value.length < c) {
                                        d.setTip(e, false, 'length', l_e ? '长度必须为' + c + '-' + l_e + '位' : '长度必须大于' + (c - 1) + '位')
                                    } else if (e.value.length > l_e) {
                                        d.setTip(e, false, 'length', c ? '长度必须为' + c + '-' + l_e + '位' : '长度必须小于' + (l_e + 1) + '位')
                                    } else if (e.getAttribute('required') == null || e.value.length > 0) {
                                        d.removeTip(e)
                                    }
                                }
                            })(this, g[i]);
                            j[j.length] = k
                        }
                    }
                    var l = (function (a) {
                        return function () {
                            for (var i = 0; i < a.length; i++) {
                                if (a[i]) {
                                    a[i].apply(this, arguments)
                                }
                            }
                        }
                    })(j);
                    g[i].onblur = l
                }
            }, validate: function (a) {
                var b;
                if (a) {
                    b = jr.dom.getsByClass(document.getElementById(a), 'ui-validate')
                } else {
                    b = jr.dom.getsByClass(document, 'ui-validate')
                }
                var c = function (e) {
                    return e.getAttribute('required') == "true" || e.getAttribute('isrequired') == "true" || e.getAttribute('length') || e.getAttribute('regex')
                };
                for (var i = 0; i < b.length; i++) {
                    if (c(b[i])) {
                        if (b[i].onblur) {
                            b[i].onblur()
                        }
                    }
                }
                return this.result(a)
            }
        }
    });
    jr.event.add(window, 'load', function () {
        jr.validator.init()
    });
});