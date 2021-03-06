/***
 * Maracatronics Robotics
 * Federal University of Pernambuco (UFPE) at Recife
 * http://www.maracatronics.com/
 *
 * This file is part of Armorial project.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 ***/

#ifndef SSLREFEREE_H
#define SSLREFEREE_H

#include <QUdpSocket>
#include <QNetworkInterface>
#include <entity/referee/mrcreferee.h>
#include <utils/basics/color.hh>
#include <entity/referee/SSLReferee/sslgameinfo.h>

#include <entity/coachview/coachview.h>
#include <entity/coachview/mainwindow.h>

class SSLReferee : public MRCReferee {
public:
    SSLReferee(QString ipAddress = "224.5.23.1", int port = 10003);
    ~SSLReferee();
    QString name();

    // Add and get game info
    SSLGameInfo* getGameInfo(Colors::Color color);
    void addGameInfo(Colors::Color color);

    // Set ball kicked
    void setBallKicked();
private:
    // Entity inherit virtual methods
    void initialization();
    void loop();
    void finalization();

    // Network configuration
    QHostAddress _ipAddress;
    int _port;

    // UDP socket
    QUdpSocket *_socket;

    // Game infos (two team colors)
    QMap<Colors::Color, SSLGameInfo*> _gameInfo;

};

#endif // SSLREFEREE_H
